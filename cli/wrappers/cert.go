package wrappers

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubectl/pkg/scheme"
	crk8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// MkcertCLI defines the methods to interact with the mkcert CLI.
type MkcertCLI interface {
	Install() error
	CARoot() (string, error)
	Certificate() string
	Key() string
}

// Mkcert represents the result of running the mkcert install command (i.e. the paths to the generated cert and key files).
type Mkcert struct {
	certPath string
	keyPath  string
}

// Install runs the mkcert install command to install the CA root certificate and key.
func (m *Mkcert) Install() error {
	if err := exec.Command("mkcert", "-install").Run(); err != nil {
		return fmt.Errorf("failed to install mkcert: %w", err)
	}
	caRoot, err := m.CARoot()
	if err != nil {
		return err
	}

	m.certPath = filepath.Join(caRoot, "rootCA.pem")
	m.keyPath = filepath.Join(caRoot, "rootCA-key.pem")

	return nil
}

// CARoot returns the path to the mkcert CA root.
func (m *Mkcert) CARoot() (string, error) {
	caRootBytes, err := exec.Command("mkcert", "-CAROOT").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get mkcert CA root: %w", err)
	}
	return string(bytes.Trim(caRootBytes, "\n")), nil
}

// Certificate returns the path to the mkcert certificate.
func (m *Mkcert) Certificate() string {
	return m.certPath
}

// Key returns the path to the mkcert key.
func (m *Mkcert) Key() string {
	return m.keyPath
}

type CACert struct {
	crk8sClient crk8sclient.Client
	mkcert      MkcertCLI
}

// NewCACert creates a new CACert instance whose purpose is to interact with
// kubernetes through a controller-runtime client instance in order to manage cert-manager resources.
func NewCACert(k8sClient K8sCLI, mkcert MkcertCLI) (*CACert, error) {
	err := certmanagerv1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, fmt.Errorf("failed to add cert-manager to scheme: %w", err)
	}

	crk8sClient, err := k8sClient.ControllerRuntimeClient(crk8sclient.Options{Scheme: scheme.Scheme})
	if err != nil {
		return nil, err
	}

	return &CACert{
		crk8sClient: crk8sClient,
		mkcert:      mkcert,
	}, nil
}

// EnsureCertManagerSecret ensures that a secret for cert-manager exists in the cluster with the given name and namespace.
// Returns a boolean indicating whether the secret already exists and an error if one occurred.
func (c *CACert) EnsureCertManagerSecret(ctx context.Context, secretName, secretNamespace string) (bool, error) {
	existingSecret := &corev1.Secret{}
	err := c.crk8sClient.Get(ctx, crk8sclient.ObjectKey{Name: secretName, Namespace: secretNamespace}, existingSecret)
	if err == nil || k8serrors.IsAlreadyExists(err) {
		return true, nil
	}
	if !k8serrors.IsNotFound(err) {
		return false, fmt.Errorf("failed to get secret: %w", err)
	}

	if err := c.mkcert.Install(); err != nil {
		return false, err
	}

	certPathBytes, err := os.ReadFile(c.mkcert.Certificate())
	if err != nil {
		return false, fmt.Errorf("failed to read cert file: %w", err)
	}

	keyPathBytes, err := os.ReadFile(c.mkcert.Key())
	if err != nil {
		return false, fmt.Errorf("failed to read key file: %w", err)
	}

	newSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: secretNamespace,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			"tls.crt": certPathBytes,
			"tls.key": keyPathBytes,
		},
	}

	if err := c.crk8sClient.Create(ctx, newSecret); err != nil {
		return false, fmt.Errorf("failed to create secret: %w", err)
	}

	return false, nil
}

// EnsureCAClusterIssuer ensures that a ClusterIssuer for cert-manager exists in the cluster with the given name and namespace.
func (c *CACert) EnsureCAClusterIssuer(ctx context.Context, secretName, clusterIssuerName, clusterIssuerNamespace string) error {
	desiredClusterIssuer := &certmanagerv1.ClusterIssuer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterIssuerName,
			Namespace: clusterIssuerNamespace,
		},
		Spec: certmanagerv1.IssuerSpec{
			IssuerConfig: certmanagerv1.IssuerConfig{
				CA: &certmanagerv1.CAIssuer{
					SecretName: secretName,
				},
			},
		},
	}

	existingClusterIssuer := &certmanagerv1.ClusterIssuer{}
	err := c.crk8sClient.Get(ctx, crk8sclient.ObjectKey{Name: clusterIssuerName}, existingClusterIssuer)
	if k8serrors.IsNotFound(err) {
		if err := c.crk8sClient.Create(ctx, desiredClusterIssuer); err != nil {
			return fmt.Errorf("failed to create ClusterIssuer: %w", err)
		}
		return nil
	}

	if err != nil && !k8serrors.IsAlreadyExists(err) {
		return fmt.Errorf("failed to get ClusterIssuer: %w", err)
	}

	if !reflect.DeepEqual(existingClusterIssuer.Spec, desiredClusterIssuer.Spec) {
		existingClusterIssuer.Spec = desiredClusterIssuer.Spec
		if err := c.crk8sClient.Update(ctx, existingClusterIssuer); err != nil {
			return fmt.Errorf("failed to update ClusterIssuer: %w", err)
		}
	}

	return nil
}

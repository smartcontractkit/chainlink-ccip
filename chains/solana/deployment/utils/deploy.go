package utils

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func MaybeDeployContract(
	b operations.Bundle,
	chain cldf_solana.Chain,
	input []datastore.AddressRef,
	contractType cldf_deployment.ContractType,
	contractVersion *semver.Version,
	contractQualifier string,
	programName string,
	programSize int,
) (datastore.AddressRef, error) {
	for _, ref := range input {
		if ref.Type == datastore.ContractType(contractType) &&
			ref.Version.Equal(contractVersion) {
			if contractQualifier != "" {
				if ref.Qualifier == contractQualifier {
					fmt.Println("Using existing", contractType, "at", ref.Address)
					return ref, nil
				}
			} else {
				fmt.Println("Using existing", contractType, "at", ref.Address)
				return ref, nil
			}
		}
	}
	programID, err := chain.DeployProgram(b.Logger, cldf_solana.ProgramInfo{
		Name:  programName,
		Bytes: programSize,
	}, false, true)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	// validate deployed programID
	_ = solana.MustPublicKeyFromBase58(programID)
	fmt.Println("Deployed", contractType, "at", programID)

	return datastore.AddressRef{
		Address:       programID,
		ChainSelector: chain.Selector,
		Type:          datastore.ContractType(contractType),
		Version:       contractVersion,
	}, nil
}

func DownloadTarGzReleaseAssetFromGithub(
	ctx context.Context,
	owner string,
	repo string,
	name string,
	tag string,
	cb func(r *tar.Reader, h *tar.Header) error,
) error {
	url := fmt.Sprintf(
		"https://github.com/%s/%s/releases/download/%s/%s",
		owner,
		repo,
		tag,
		name,
	)

	_, err := withGetRequest(ctx, url, func(res *http.Response) (any, error) {
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("request failed with status %d - could not download tar.gz release artifact from Github (url = '%s')", res.StatusCode, url)
		}

		gzipReader, err := gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer gzipReader.Close()

		tarReader := tar.NewReader(gzipReader)
		for {
			header, err := tarReader.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			if err := cb(tarReader, header); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}

func DownloadSolanaCCIPProgramArtifacts(ctx context.Context, dir string, sha string) error {
	const ownr = "smartcontractkit"
	const repo = "chainlink-ccip"
	const name = "artifacts.tar.gz"
	tag := "solana-artifacts-localtest-" + sha

	return DownloadTarGzReleaseAssetFromGithub(ctx, ownr, repo, name, tag, func(r *tar.Reader, h *tar.Header) error {
		if h.Typeflag != tar.TypeReg {
			return nil
		}

		outPath := filepath.Join(dir, filepath.Base(h.Name))
		if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, r); err != nil {
			return err
		}

		return nil
	})
}

func withGetRequest[T any](ctx context.Context, url string, cb func(res *http.Response) (T, error)) (T, error) {
	var empty T

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return empty, err
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()

	return cb(res)
}

package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func TestManualRegistration_VerifyPreconditions(t *testing.T) {
	const (
		chainA = uint64(5009297550715157269)
		chainB = uint64(15971525489660198786)
	)

	baseRegistration := func() tokens.RegisterTokenConfig {
		return tokens.RegisterTokenConfig{
			ChainSelector: chainA,
			TokenPoolRef: datastore.AddressRef{
				Type:          "TokenPool",
				Version:       semver.MustParse("1.0.0"),
				ChainSelector: chainA,
				Qualifier:     "default",
			},
			ProposedOwner: "0x0000000000000000000000000000000000000001",
		}
	}

	tests := []struct {
		name        string
		cfg         tokens.ManualRegistrationInput
		expectedErr string
	}{
		{
			name:        "Failure - no registrations",
			cfg:         tokens.ManualRegistrationInput{},
			expectedErr: "at least one registration is required",
		},
		{
			name: "Success - valid registration",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{baseRegistration()},
			},
		},
		{
			name: "Failure - invalid chain selector",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{
					func() tokens.RegisterTokenConfig {
						registration := baseRegistration()
						registration.ChainSelector = 0
						return registration
					}(),
				},
			},
			expectedErr: "invalid chain selector 0",
		},
		{
			name: "Failure - duplicate chain selector",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{baseRegistration(), baseRegistration()},
			},
			expectedErr: "duplicate entry for chain selector",
		},
		{
			name: "Failure - missing ProposedOwner",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{
					func() tokens.RegisterTokenConfig {
						registration := baseRegistration()
						registration.ProposedOwner = ""
						return registration
					}(),
				},
			},
			expectedErr: "ProposedOwner is required",
		},
		{
			name: "Failure - both refs empty",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{
					func() tokens.RegisterTokenConfig {
						registration := baseRegistration()
						registration.TokenPoolRef = datastore.AddressRef{}
						return registration
					}(),
				},
			},
			expectedErr: "at least one of TokenPoolRef or TokenRef is required",
		},
		{
			name: "Failure - TokenPoolRef chain selector mismatch",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{
					func() tokens.RegisterTokenConfig {
						registration := baseRegistration()
						registration.TokenPoolRef.ChainSelector = chainB
						return registration
					}(),
				},
			},
			expectedErr: "TokenPoolRef.ChainSelector",
		},
		{
			name: "Failure - TokenRef chain selector mismatch",
			cfg: tokens.ManualRegistrationInput{
				Registrations: []tokens.RegisterTokenConfig{
					func() tokens.RegisterTokenConfig {
						registration := baseRegistration()
						registration.TokenRef = datastore.AddressRef{
							Type:          "Token",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: chainB,
						}
						return registration
					}(),
				},
			},
			expectedErr: "TokenRef.ChainSelector",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			changeset := tokens.ManualRegistration()
			err := changeset.VerifyPreconditions(deployment.Environment{}, tc.cfg)

			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}

package scripts

import "testing"

func TestDeployHomeChain(t *testing.T) {
	t.Skip()
	t.Parallel()
	env := TestEnvFWOG()

	DeployHomeChain(logger, env)
}

package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
)

func TestMain(m *testing.M) {
	// to remove containers after the tests automatically
	_ = os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "false")
	// to isolate containers the same way we do in e2e environment
	err := framework.DefaultNetwork(nil)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

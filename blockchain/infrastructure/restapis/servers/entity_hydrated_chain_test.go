package servers

import (
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_chain_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a chain:
	chain := chains.CreateChainForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, chain, t)
}

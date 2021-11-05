package test

import (
	"testing"

	"github.com/iotaledger/wasp/alpha-interface/src/contract"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, alphainterfacecontract.ScName, alphainterfacecontract.OnLoad)
	require.NoError(t, ctx.ContractExists(alphainterfacecontract.ScName))
}

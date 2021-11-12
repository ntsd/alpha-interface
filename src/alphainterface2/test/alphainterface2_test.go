package test

import (
	"testing"

	"github.com/iotaledger/wasp/alphainterface/src/alphainterface2/go/alphainterface2"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, alphainterface2.ScName, alphainterface2.OnLoad)
	require.NoError(t, ctx.ContractExists(alphainterface2.ScName))
}

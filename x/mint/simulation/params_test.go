package simulation_test

import (
	//#nosec
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sge-network/sge/x/mint/simulation"
)

func TestParamChangest(t *testing.T) {
	s := rand.NewSource(1)
	//#nosec
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"mint/BlocksPerYear", "BlocksPerYear", "\"6311520.000000000000000000\"", "mint"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 1)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}

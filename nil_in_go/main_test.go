package nil_in_go

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNil(t *testing.T) {
	var err error
	var x any
	x = err
	require.True(t, x == nil)
}

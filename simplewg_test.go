package simplewg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSimpleWg tests the basic functionality of the Wg type from the simplewg package.
func TestSimpleWg(t *testing.T) {
	var wg Wg
	var i int
	// Test that Go executes the function and sets i to 123.
	require.True(t, wg.Go(func() { i = 123 }))
	wg.Wait()
	require.Equal(t, 123, i)
	// Test that subsequent calls to Go return false after Wait has been called.
	require.False(t, wg.Go(func() {}))
}

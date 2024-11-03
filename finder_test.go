package gomodfinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	t.Run("failed with 0 levels", func(t *testing.T) {
		_, err := Find("./", 0)

		assert.Equal(t, &FileNotFoundError{
			File:      "go.mod",
			Locations: []string{},
		}, err)
	})

	t.Run("success", func(t *testing.T) {
		mf, err := Find("./", 1)
		require.NoError(t, err)

		assert.Equal(t, "github.com/artarts36/gomodfinder", mf.Module.Mod.Path)
	})
}

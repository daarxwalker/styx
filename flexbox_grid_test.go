package styx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlexboxGrid(t *testing.T) {
	t.Run(
		"grid three same cols", func(t *testing.T) {
			styler := New()
			gridColsStyle := styler.Build().
				Grid().GridCols(3)
			assert.Equal(
				t,
				`grid grid-cols-3`,
				gridColsStyle.Class(),
			)
			assert.Equal(
				t,
				`.grid{display:grid;}.grid-cols-3{grid-template-columns:repeat(3,minmax(0,1fr));}`,
				styler.String(),
			)
		},
	)
}

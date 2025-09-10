package styx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	t.Run(
		"negative rotate", func(t *testing.T) {
			styler := New()
			rotateStyle := styler.Build().
				Rotate("-180deg")
			assert.Equal(
				t,
				`-rotate-180deg`,
				rotateStyle.Class(),
			)
			assert.Equal(
				t,
				`.-rotate-180deg{rotate:-180deg;}`,
				styler.String(),
			)
		},
	)
}

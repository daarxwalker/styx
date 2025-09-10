package styx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseStyles(t *testing.T) {
	t.Run(
		"create base styles", func(t *testing.T) {
			instance := New()
			instance.CreateBaseStyles()
			assert.Equal(
				t,
				true,
				len(instance.String()) > 0,
			)
			assert.Equal(
				t,
				true,
				strings.HasPrefix(instance.String(), ":root"),
			)
		},
	)
}

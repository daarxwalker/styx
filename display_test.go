package styx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	t.Run(
		"flex", func(t *testing.T) {
			instance := New().Build().Flex(Hover, Lg)
			assert.Equal(
				t,
				`hover:lg:flex`,
				instance.Class(),
			)
			assert.Equal(
				t,
				fmt.Sprintf(`@media (min-width:%s) {.hover\:lg\:flex{&:hover{display:flex;}}}`, BaseBreakpoints[Lg]),
				instance.Style(),
			)
		},
	)
}

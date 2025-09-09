package styx

import (
	"fmt"
	"testing"
)

func TestBaseStyles(t *testing.T) {
	t.Run(
		"create base styles", func(t *testing.T) {
			instance := New()
			instance.CreateBaseStyles()
			fmt.Println(instance.String())
			// assert.Equal(
			// 	t,
			// 	`lg:inline-flex bg-success px-1rem py-0.5rem`,
			// 	button.Class(),
			// )
		},
	)
}

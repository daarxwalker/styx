
# STYX
> Type-safe, composable, utility-first CSS engine in Go
```go  
package app  
  
import (
	. "github.com/daarxwalker/gox"
	"github.com/daarxwalker/styx"
)  
  
var styler = styx.New()  
  
var (  
	someStyle = styler.Build().
		Relative().
		BgColor(ColorWhite).
		PX("1rem").PY("0.5rem").
		Border("1px").BorderColor(ColorNeutral).
		Rounded("1rem").
		Class()
)  
  
func SomeComponent(nodes ...Node) Node {  
	return Div(
		Class(someStyle)
		Fragment(nodes...),
	)
}  
  
```
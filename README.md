
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

<br>

## Why Choose STYX?
> If you love writing Go and don’t want to rely on npm packages just to build reliable, scalable styles.

Styx brings proven utility-first principles — like those from TailwindCSS — into the Go ecosystem, but with a type-safe, composable, and idiomatic API. Design styles directly in Go, avoid runtime CSS pitfalls, and keep full control without leaving your language:

* **🔮🚫No magic:** What you write is exactly what you get; no hidden rules, no preprocessing, no side effects.
* **🧱 Utility-first CSS generation:** Compose styles using chainable Go methods inspired by TailwindCSS.
* **🧠 Type-safe API:** Build CSS utilities through strongly typed Go functions instead of raw strings.
* **🧵 CSS deduplication & optimization:** Prevents duplicate class definitions and style blocks automatically.
* **🎯 Composable styles:** Combine and reuse style sets (Join, Build, etc.) across components or packages.
* **🎛 Custom selectors support:** Write scoped styles with custom selectors (.button, *, body > div, etc.).
* **🎨 CSS variables & theming:** Define and use CSS custom properties (breakpoints,colors,fonts,shadows).
* **📱 Responsive modifiers & pseudo-classes:** Add media queries and support for hover, focus, disabled, and more via modifiers.
* **📦 Fully embeddable:** Integrate seamlessly in Go applications – ideal for web apps, SSR, or static generation tools.
* **🧹 Browser reset & base styles:** Generate reset styles and global base styles programmatically.
* **🪶 No external dependencies:** Pure Go implementation with zero runtime dependencies.

<br>

## Installation
```bash
go get github.com/daarxwalker/styx
```

<br>

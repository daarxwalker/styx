
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

## Features
- [Scope](#scope)
- [Build](#build)
- [Modifier](#modifier)
- [Custom selector](#custom-selector)
- [Define CSS variables](#define-css-variables)
- [Join](#join)
- [Result](#result)

<br>

### Scope
Create multiple Styx instances<br>
> _Useful for different apps, modules, services, etc..._
```go
instance01 := styx.New()
instance02 := styx.New()
```

<br>

### Build
Begin write style with .Build() method.<br>
> _Method accepts custom selector argument, which group all defined props to single style._
```go
var (
	styler := styx.New()
)

var (
	someCSS := styler.Build().
				Relative().
				PX("1rem", Lg).
				TextCenter()
)

func main() {
	someCSS.Class() // Use in GOX (HTML)
	someCSS.Style() // Useful for tests
}
```

<br>

### Modifier
All props have optional Modifier type arguments<br>
> _Useful for breakpoints, pseudo-selectors, pseudo-elements, theming._
```go
var (
	styler := styx.New()
)

var (
    someCSS := styler.Build().
            Width("100%", Lg).
            BgColor(ColorNeutral).
            BgColor(ColorNeutralHover, Hover)
)

```

<br>

### Custom selector
Define custom selector in .Build(".custom-class") method.<br>
> _It groups all defined props to single style._
> _Finish, Style or Class method is required to close group._
```go
var (
	styler := styx.New()
)

var (
	styler.Build(".some-custom-class").
		Relative().
		P("1rem").
		Finish()
)

func main() {
	fmt.Printlin(styler.String()) // .some-custom-class{position:relative;padding:1rem;}
}
```

<br>

### Define CSS variables
Join multiple Styx instances<br>
> _Useful for UI libraries, base styles._
```go
var (
	styler := styx.New()
)

var (
	styler.Build(".some-custom-class").
		DefineVar("name", "value").
		Finish()
)
```

<br>

### Join
Join multiple Styx instances<br>
> _Useful for UI libraries, base styles._
```go
instance01 := styx.New()
instance02 := styx.New()
instance01.Join(instance02)
```
<br>

### Result
Retrieve the final styles in an idiomatic, Go-centric way.<br>
> _Useful for writing styles directly to files, HTTP responses, or embedded assets._
```go
stylesBytes := styler.Bytes() // Get CSS as []byte
```
```go
stylesString := styler.String() // Get CSS as string
```
```go
writeErr := styler.Write(writer) // Write CSS to io.Writer
```

<br>
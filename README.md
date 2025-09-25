
# STYX
> CSS in Go, no magic, no Node
```go  
package app  
  
import (
	. "github.com/daarxwalker/gox"
	"github.com/daarxwalker/styx"
)  
  
var styler = styx.New()  
  
var (  
    someButtonClass = styler.Rule(
	    styx.WithSelector(styx.Class("btn")),
	    styx.WithProp(styx.Padding, styx.Rem(1)),
	    styx.WithProp(styx.BackgroundColor, "blue"),
	    styx.WithRule(
		    styx.WithSelector(styx.Self+styx.Hover),
		    styx.WithProp(styx.BackgroundColor, "red"),
	    ),
    )
)  
  
func SomeButtonComponent(nodes ...Node) Node {  
	return Button(
		Class(someButtonClass),
		Type("button"),
		Fragment(nodes...),
	)
}  
  
```

<br>

## When to use
> Use Styx if you want type-safe CSS in Go without depending on Node.js or if you’re building Go-first SSR apps, micro-frontends, or design systems.

* **🔮🚫No magic:** What you write is exactly what you get; no hidden rules, no preprocessing, no side effects.
* **🧠 Type-safe API:** Build CSS utilities through strongly typed Go functions instead of raw strings.
* **🎯 Composable styles:** Combine and reuse style sets (Join, Rule, Fragment) across components or packages.
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
- [Join](#join)
- [Rule](#rule)
- [Prop](#prop)
- [Fragment](#fragment)
- [Nested rule](#nested-rule)
- [Global rule](#global-rule)
- [Helpers chaining](#helpers-chaining)
- [Result](#result)

<br>

### Scope
Create multiple Styx instances<br>
> _Useful for different apps, modules, services, etc..._
```go
styler01 := styx.New()
styler02 := styx.New()
```

<br>

### Join
Join multiple Styx instances<br>
> _Useful for UI libraries, base styles._
```go
styler01 := styx.New()
styler02 := styx.New()
styler01.Join(styler02)
```
<br>

### Rule
Begin to write statement with .Rule() method.<br>
> _Method accepts options pattern functions like WithSelector(), WithProp(), WithRule()._
```go
var (
    styler := styx.New()
)

func createStyles() {
    styler.Rule(
      styx.WithSelector(styx.Class("example")),
      styx.WithProp(styx.Background, "red"),
    )
}

func main() {
    fmt.Println(styler.String()) // .example{background:red;}
}
```

<br>

### Prop
Define CSS prop<br>
> _Method accepts prop name and variadic value._
```go
styx.WithProp(styx.Border, Px(1), OutlineSolid, "black")
```

<br>

### Fragment
Props can be grouped in flat way.<br>
> _Useful for reusable props._
```go
var (
    styler := styx.New()
)

func reusableStyles() styx.Node {
    return styx.Fragment(
        styx.WithProp(styx.Padding, styx.Rem(1)),
    )
}

func createStyles() {
    styler.Rule(
        styx.WithSelector(styx.Class("example")),
        styx.WithProp(styx.Background, "red"),
        reusableStyles(),
    )
}

```

<br>

### Nested rule
Define nested rule, like child, pseudo-classes, pseudo-elements,at-rules.<br>
> _Useful for :hover, :focus, @media, etc..._
```go
var (
	styler := styx.New()
)

func createStyles() {
    styler.Rule(
        styx.WithSelector(styx.Class("example")),
        styx.WithProp(styx.Background, "red"),
        styx.WithRule(
            styx.WithSelector(styx.Self+styx.Hover),
            styx.WithProp(styx.Background, "blue"),
        )
        // Selector helpers chaining
        styx.WithRule(
            styx.WithSelector(styx.Self+styx.Hover),
            styx.WithProp(styx.Border, "blue"),
        )
    )
}
```

<br>

### Global rule
Create rules in global scope.<br>
> _Define global rules, like :root, @font-face._
```go
var (
    styler := styx.New()
)

func createGlobalStyles() {
    styler.Global(
        styx.WithRule(
            styx.WithSelector(styx.Root),
            styx.DefineVar(fontSize, "16px"),
            styx.DefineVar(fontFamily, "Inter,sans-serif"),
        )
    )
}
```

<br>

### Helpers chaining
Define selector with helpers.<br>
> _Useful for special type-safe selectors._

> Self, Child, NextSibling, Sibling, Class(), Id(), Attribute()
```go
var (
	styler := styx.New()
)

func createStyles() {
    styler.Rule(
        styx.WithSelector(styx.Class("example")),
        styx.WithRule(
            styx.WithSelector(styx.Self, styx.Child, styx.Class('some-child')),
        )
    )
}
```

<br>

### Helpers
Sugared rules<br>
> _Useful for defining @font-face._
```go
var (
    styler := styx.New()
)

func createGlobalStyles() {
    // FontFace()
    styler.FontFace(
        "Inter",
        "100 900",
        "normal",
        "/fonts/Inter-VariableFont_opsz,wght.ttf",
        "truetype-variations",
    )
		
    // ResetStyles()
    styler.ResetStyles()
}
```

<br>

### Result
Retrieve the final styles in an idiomatic, Go-centric way.<br>
> _Styx generates CSS you write, no big runtimes, no processing._
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
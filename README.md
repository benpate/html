# builder 🚧

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/benpate/builder)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/builder?style=flat-square)](https://goreportcard.com/report/github.com/benpate/builder)
[![Build Status](http://img.shields.io/travis/benpate/builder.svg?style=flat-square)](https://travis-ci.org/benpate/builder)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/builder.svg?style=flat-square)](https://codecov.io/gh/benpate/builder)
![Version](https://img.shields.io/github/v/release/benpate/builder?include_prereleases&style=flat-square&color=brightgreen)

## Efficient HTML Tag Assembly

Builder is a simple library for assembling HTML tags using Go.  

```go
b := builder.New()

b.Div().Class("wrapper")
b.Div().Class("inner")
b.Form().Attr("action", "my-server")
b.Input().Name("FullName").Value("John Connor").Close()
b.Input().Name("Email").Value("john@connor.mil").Close()
b.CloseAll()
b.String()
```

## Why Builder?

Why not just use [Go Templates](https://golang.org/pkg/text/template/) instead?  Templates work great in many cases, but they can be cumbersome when building complex conditional logic.  Builder uses an efficient strings.Builder](https://pkg.go.dev/strings#Builder) to assemble the exact HTML you need, and nothing extra.


## Pull Requests Welcome

This library is growing rapidly, as the requirements of its downstram projects continue to evolve.  How can it help you build your next masterpiece?  Add your voice, because we're all in this together! 🚧
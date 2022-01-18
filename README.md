# html 🚧

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/html)
[![Build Status](https://img.shields.io/github/workflow/status/benpate/html/Go/master)](https://github.com/benpate/html/actions/workflows/go.yml)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/html.svg?style=flat-square)](https://codecov.io/gh/benpate/html)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/html?style=flat-square)](https://goreportcard.com/report/github.com/benpate/html)
[![Version](https://img.shields.io/github/v/release/benpate/html?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/html/releases)

## Efficient HTML Tag Assembly

This library builds a string buffer of HTML tags programmatically.

```go
b := html.New()

b.Div().Class("wrapper")
b.Div().Class("inner")
b.Form().Attr("action", "my-server")
b.Input().Name("FullName").Value("John Connor").Close()
b.Input().Name("Email").Value("john@connor.mil").Close()
b.CloseAll()
b.String()
```

## Why Builder?

Why not just use [Go Templates](https://golang.org/pkg/text/template/) instead?  Templates work great in many cases, but they can be cumbersome when building complex conditional logic directly in your code.  Builder uses an efficient [strings.Builder](https://pkg.go.dev/strings#Builder) to assemble the exact HTML you need, and nothing extra.

## Pull Requests Welcome

This library is growing rapidly, as the requirements of its downstram projects continue to evolve.  How can it help you build your next masterpiece?  Add your voice, because we're all in this together! 🚧

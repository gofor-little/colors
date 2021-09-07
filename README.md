## A package for ANSI color formatting

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/colors?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/colors)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/colors/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gofor-little/colors/CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/colors)](https://goreportcard.com/report/github.com/gofor-little/colors)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/colors)](https://pkg.go.dev/github.com/gofor-little/colors)

### Introduction
* ANSI color formatting for consoles
* No dependencies outside the standard library

### Example
```go
package main

import (
	"fmt"

	"github.com/gofor-little/colors"
)

func main() {
	fmt.Println(colors.Black("Black"))
	fmt.Println(colors.BlackBright("BlackBright"))
	fmt.Println(colors.Red("Red"))
	fmt.Println(colors.RedBright("RedBright"))
	fmt.Println(colors.Green("Green"))
	fmt.Println(colors.GreenBright("GreenBright"))
	fmt.Println(colors.Yellow("Yellow"))
	fmt.Println(colors.YellowBright("YellowBright"))
	fmt.Println(colors.Blue("Blue"))
	fmt.Println(colors.BlueBright("BlueBright"))
	fmt.Println(colors.Magenta("Magenta"))
	fmt.Println(colors.MagentaBright("MagentaBright"))
	fmt.Println(colors.Cyan("Cyan"))
	fmt.Println(colors.CyanBright("CyanBright"))
	fmt.Println(colors.White("White"))
	fmt.Println(colors.WhiteBright("WhiteBright"))
}
```
![](https://github.com/gofor-little/colors/blob/main/example-dark.png)
![](https://github.com/gofor-little/colors/blob/main/example-light.png)

### Testing
Run ```go test ./...``` in the root directory.

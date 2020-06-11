# little colors

[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/colors)](https://goreportcard.com/report/github.com/gofor-little/colors)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofor-little/colors)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/v/release/gofor-little/colors?include_prereleases)](https://github.com/gofor-little/colors/releases )

## Example
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

## Output
![](https://github.com/gofor-little/colors/blob/master/example-dark.png)
![](https://github.com/gofor-little/colors/blob/master/example-light.png)
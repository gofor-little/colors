package colors_test

import (
	"fmt"
	"testing"

	"github.com/gofor-little/colors"
)

func TestColors(t *testing.T) {
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

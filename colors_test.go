package colors_test

import (
	"fmt"
	"testing"

	"github.com/gofor-little/colors"
)

func TestColors(t *testing.T) {
	testCases := []struct {
		name  string
		color func(...interface{}) string
		s     string
		want  string
	}{
		{name: "Test_colors.Black()", color: colors.Black, s: "Black", want: fmt.Sprintf("\u001B[1;30m%s\u001B[0m", "Black")},
		{name: "Test_colors.BlackBright()", color: colors.BlackBright, s: "BlackBright", want: fmt.Sprintf("\u001B[1;90m%s\u001B[0m", "BlackBright")},
		{name: "Test_colors.Red()", color: colors.Red, s: "Red", want: fmt.Sprintf("\u001B[1;31m%s\u001B[0m", "Red")},
		{name: "Test_colors.RedBright()", color: colors.RedBright, s: "RedBright", want: fmt.Sprintf("\u001B[1;91m%s\u001B[0m", "RedBright")},
		{name: "Test_colors.Green()", color: colors.Green, s: "Green", want: fmt.Sprintf("\u001B[1;32m%s\u001B[0m", "Green")},
		{name: "Test_colors.GreenBright()", color: colors.GreenBright, s: "GreenBright", want: fmt.Sprintf("\u001B[1;92m%s\u001B[0m", "GreenBright")},
		{name: "Test_colors.Yellow()", color: colors.Yellow, s: "Yellow", want: fmt.Sprintf("\u001B[1;33m%s\u001B[0m", "Yellow")},
		{name: "Test_colors.YellowBright()", color: colors.YellowBright, s: "YellowBright", want: fmt.Sprintf("\u001B[1;93m%s\u001B[0m", "YellowBright")},
		{name: "Test_colors.Blue()", color: colors.Blue, s: "Blue", want: fmt.Sprintf("\u001B[1;34m%s\u001B[0m", "Blue")},
		{name: "Test_colors.BlueBright()", color: colors.BlueBright, s: "BlueBright", want: fmt.Sprintf("\u001B[1;94m%s\u001B[0m", "BlueBright")},
		{name: "Test_colors.Magenta()", color: colors.Magenta, s: "Magenta", want: fmt.Sprintf("\u001B[1;35m%s\u001B[0m", "Magenta")},
		{name: "Test_colors.MagentaBright()", color: colors.MagentaBright, s: "MagentaBright", want: fmt.Sprintf("\u001B[1;95m%s\u001B[0m", "MagentaBright")},
		{name: "Test_colors.Cyan()", color: colors.Cyan, s: "Cyan", want: fmt.Sprintf("\u001B[1;36m%s\u001B[0m", "Cyan")},
		{name: "Test_colors.CyanBright()", color: colors.CyanBright, s: "CyanBright", want: fmt.Sprintf("\u001B[1;96m%s\u001B[0m", "CyanBright")},
		{name: "Test_colors.White()", color: colors.White, s: "White", want: fmt.Sprintf("\u001B[1;37m%s\u001B[0m", "White")},
		{name: "Test_colors.WhiteBright()", color: colors.WhiteBright, s: "WhiteBright", want: fmt.Sprintf("\u001B[1;97m%s\u001B[0m", "WhiteBright")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.color(tc.s); got != tc.want {
				t.Fatalf("unexpected result returned from string: %s, wanted: %s, got: %s", tc.s, tc.want, got)
			}
		})
	}
}

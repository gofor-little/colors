package colors

import "fmt"

// Available colors.
var (
	Black         = Color("\033[1;30m%s\033[0m")
	BlackBright   = Color("\033[1;90m%s\033[0m")
	Red           = Color("\033[1;31m%s\033[0m")
	RedBright     = Color("\033[1;91m%s\033[0m")
	Green         = Color("\033[1;32m%s\033[0m")
	GreenBright   = Color("\033[1;92m%s\033[0m")
	Yellow        = Color("\033[1;33m%s\033[0m")
	YellowBright  = Color("\033[1;93m%s\033[0m")
	Blue          = Color("\033[1;34m%s\033[0m")
	BlueBright    = Color("\033[1;94m%s\033[0m")
	Magenta       = Color("\033[1;35m%s\033[0m")
	MagentaBright = Color("\033[1;95m%s\033[0m")
	Cyan          = Color("\033[1;36m%s\033[0m")
	CyanBright    = Color("\033[1;96m%s\033[0m")
	White         = Color("\033[1;37m%s\033[0m")
	WhiteBright   = Color("\033[1;97m%s\033[0m")
)

// Color formats the ANSI color code with the passed args.
func Color(ansiColorCode string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(ansiColorCode, args...)
	}
}

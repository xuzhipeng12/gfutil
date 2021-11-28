/**
 * @Author xuzhipeng
 * @Description
 * @Date 2021/11/27 11:19 下午
 **/
package tools

import (
	"fmt"
)

var (
	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

func SprintlnDirectory(a ...interface{}) string {
	return fmt.Sprintln(strAppend(blue, a...)...)
}

func SprintDirectory(a ...interface{}) string {
	return fmt.Sprint(strAppend(blue, a...)...)
}

func SprintfDirectory(format string, a ...interface{}) string {
	return fmt.Sprintf(format, strAppend(blue, a...)...)
}

func SprintlnFile(a ...interface{}) string {
	return fmt.Sprintln(strAppend(green, a...)...)
}

func SprintFile(a ...interface{}) string {
	return fmt.Sprint(strAppend(green, a...)...)
}
func SprintfFile(format string, a ...interface{}) string {
	return fmt.Sprintf(format, strAppend(green, a...)...)
}

func strAppend(color string, a ...interface{}) []interface{} {
	b := make([]interface{}, 0)
	b = append(b, color)
	b = append(b, a...)
	b = append(b, reset)
	return b
}

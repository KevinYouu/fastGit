package colors

// COLORS
var COLORS = map[string]string{
	"black":          "\x1b[30m",
	"red":            "\x1b[31m",
	"green":          "\x1b[32m",
	"yellow":         "\x1b[33m",
	"blue":           "\x1b[34m",
	"magenta":        "\x1b[35m",
	"cyan":           "\x1b[36m",
	"white":          "\x1b[37m",
	"gray":           "\x1b[90m",
	"reset":          "\x1b[0m",
	"bold":           "\x1b[1m",
	"underline":      "\x1b[4m",
	"blink":          "\x1b[5m",
	"reverse":        "\x1b[7m",
	"black_weak":     "\x1b[30;1m",
	"red_weak":       "\x1b[31;1m",
	"green_weak":     "\x1b[32;1m",
	"yellow_weak":    "\x1b[33;1m",
	"blue_weak":      "\x1b[34;1m",
	"magenta_weak":   "\x1b[35;1m",
	"cyan_weak":      "\x1b[36;1m",
	"white_weak":     "\x1b[37;1m",
	"gray_weak":      "\x1b[90;1m",
	"reset_weak":     "\x1b[0m",
	"bold_weak":      "\x1b[1;1m",
	"underline_weak": "\x1b[4;1m",
	"blink_weak":     "\x1b[5;1m",
	"reverse_weak":   "\x1b[7;1m",
}

// RenderColor should be used to render colored text
func RenderColor(color string, text string) string {
	code, ok := COLORS[color]
	if !ok {
		return text
	}
	return code + text + "\x1b[0m"
}

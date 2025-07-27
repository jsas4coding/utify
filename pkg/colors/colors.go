package colors

const (
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Cyan      = "\033[36m"
	White     = "\033[97m"
	Magenta   = "\033[35m"
	Gray      = "\033[90m"
	LightBlue = "\033[94m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Reset     = "\033[0m"
)

var userColors = map[string]string{}

func GetUserColor(key string) (string, bool) {
	color, exists := userColors[key]
	return color, exists
}

func SetColorTable(newColors map[string]string) {
	for k, v := range newColors {
		userColors[k] = v
	}
}

func ClearUserColors() {
	userColors = make(map[string]string)
}
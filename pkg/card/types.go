package card

type Theme int

const (
	Default Theme = iota
	Dark
)

type Svg interface {
	Svg() string
}

type Card interface {
	AddContent(content interface{})
	Length() int
	SwitchTheme(theme Theme)
}

type Config struct {
	Title      string
	Width      int
	LineHeight int
}

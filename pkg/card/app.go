package card

import (
	"fmt"
	"strings"
)

type Type int

const (
	List Type = iota
	Table
)

func NewListCard(config Config) *ListCard {
	return &ListCard{
		CImpl:   NewCard(config),
		Content: []string{},
	}
}

func NewTableCard(config Config, column []Column) *TableCard {
	return &TableCard{
		CImpl:   NewCard(config),
		Columns: column,
		Content: [][]string{},
	}
}

func NewCard(config Config) *CImpl {
	card := &CImpl{
		Title:      config.Title,
		Width:      config.Width,
		LineHeight: config.LineHeight,
		Theme:      Default,
		Content:    "Empty",
	}
	return card
}

type CImpl struct {
	Title      string
	Width      int
	LineHeight int
	Theme      Theme
	Content    interface{}
}

func (c *CImpl) SwitchTheme(theme Theme) {
	c.Theme = theme
}

func (c *CImpl) AddContent(content interface{}) {
	c.Content = content
}

func (c *CImpl) Length() int {
	return 1
}

// TODO: 支持一下theme
func (c *CImpl) Svg() string {
	builder := new(strings.Builder)
	height := c.LineHeight + 60
	// Svg start
	builder.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="%d" height="%d">`, c.Width, height))
	// Title
	titleHeight := 35
	builder.WriteString(fmt.Sprintf(`<g transform="translate(25, 35)"><text x="0" y="0" class="header">%s</text></g>`, c.Title))
	// Body
	builder.WriteString(fmt.Sprintf(`<g  transform="translate(25, %d)"><text x="0" y="0" class="line">%s</text></g>`, c.LineHeight+titleHeight, c.Content))
	// Svg end
	builder.WriteString(`</svg>`)

	return builder.String()
}

package card

import (
	"fmt"
	"strings"
)

type ListCard struct {
	*CImpl
	Content []string
}

func (c *ListCard) AddContent(content interface{}) {
	c.Content = append(c.Content, content.(string))
}

func (c *ListCard) Length() int {
	return len(c.Content)
}

func (c *ListCard) Svg() string {
	builder := new(strings.Builder)
	height := len(c.Content)*c.LineHeight + 60
	// Svg start
	builder.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="%d" height="%d">`, c.Width, height))
	// Css
	builder.WriteString(`<style>
.header{
fill: #fff;
font-size: 18px;
font-weight: 600;
font-family: "SourceHanSerifCN-Bold",PingFang SC,-apple-system,SF UI Text,Lucida Grande,STheiti,Microsoft YaHei,sans-serif;
}
.line{
fill:#eef0f4;
font-family: "SourceHanSerifCN-Bold",PingFang SC,-apple-system,SF UI Text,Lucida Grande,STheiti,Microsoft YaHei,sans-serif;
width: 90%;
overflow-x: hidden;
}
</style>`)
	// Bg
	builder.WriteString(`<rect xmlns="http://www.w3.org/2000/svg" x="0.5" y="0.5" rx="4.5" height="99%" stroke="#27282a" width="99%" fill="#27282a" stroke-opacity="1"/>`)
	// Title
	titleHeight := 35
	builder.WriteString(fmt.Sprintf(`<g transform="translate(25, 35)"><text x="0" y="0" class="header">%s</text></g>`, c.Title))
	for i, line := range c.Content {
		builder.WriteString(fmt.Sprintf(`<g  transform="translate(25, %d)"><text x="0" y="0" class="line">%s</text></g>
`, (i+1)*c.LineHeight+titleHeight, line))
	}
	// Svg end
	builder.WriteString(`</svg>`)

	return builder.String()
}

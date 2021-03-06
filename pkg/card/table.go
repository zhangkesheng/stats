package card

import (
	"fmt"
	"strings"
)

type Column struct {
	Title string
	Width float32
	Style string
}

type TableCard struct {
	*CImpl
	ShowHeader bool
	Columns    []Column
	Content    [][]string
}

func (c *TableCard) AddContent(content interface{}) {
	c.Content = append(c.Content, content.([]string))
}

func (c *TableCard) Length() int {
	return len(c.Content)
}

func (c *TableCard) Svg() string {
	start := 0
	if c.ShowHeader {
		start = 1
	}
	builder := new(strings.Builder)
	height := (len(c.Content)+start)*c.LineHeight + 60
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

	var lines [][]string

	if c.ShowHeader {
		var header []string
		for _, column := range c.Columns {
			header = append(header, column.Title)
		}
	}

	lines = append(lines, c.Content...)
	for i, line := range lines {
		cb := new(strings.Builder)
		cb.WriteString(fmt.Sprintf(`<g class="line" transform="translate(25, %d)">`, (i+1+start)*c.LineHeight+titleHeight))
		for j, col := range line {
			y := float32(0)
			if j > 0 {
				y = c.Columns[j-1].Width * float32(c.Width)
			}
			cb.WriteString(fmt.Sprintf(`<text y="0" x="%f" width="%f" style="%s">%s</text>`, y, c.Columns[j].Width*float32(c.Width), c.Columns[j].Style, col))
		}
		cb.WriteString("</g>\n")
		builder.WriteString(cb.String())
	}
	// Svg end
	builder.WriteString(`</svg>`)

	return builder.String()
}

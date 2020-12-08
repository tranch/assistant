package components

import (
	"fmt"
	"html/template"
)

type Text struct {
	Name  string
	Title string
}

func (c *Text) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<h2>%s</h2>`, c.Title))
}

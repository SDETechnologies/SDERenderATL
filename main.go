package main

import (
	"context"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"patrick.com/render-atl-hackathon/chatgpt"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	parsed, err := template.ParseGlob("views/*.html")
	if err != nil {
		panic(err)
	}
	return &Template{
		tmpl: parsed,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {

	review := "One of Route 185 drivers, has the nastiest personality / most rude person ever. The male driver for the morning route has the nerve to stop every morning to run into Chick fil la to get his food making people late. He only wants to eat, make rude comments, and lean on the window while he drives. He also brakes hard because he is not paying attention to the cars in front of him. He needs to be fired he simply don't care about his job. I bet when he was applying for the job he indicated he love working with people.... Not True...."

	res, err := chatgpt.SummarizeReview(context.Background(), review)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)

}

package templates

import (
	"bytes"
	"text/template"
	"log"
)

type Data struct {
	Title string
	Description string
}

type CustomDOM struct {
	buff bytes.Buffer
}

func (c *CustomDOM) Write(p []byte) (int, error) {
	c.buff.Write(p)
	log.Printf("[CustomDOM]: %q", string(p))
	return len(p), nil
}

func (c *CustomDOM) String() string {
	return c.buff.String()
}

func BuildTemplate(userTemplate string, d *Data) (string, error) {
	tmpl := template.Must(template.New("user-template").Parse(userTemplate))

	c := &CustomDOM{}
	if err := tmpl.Execute(c, d); err != nil {
		return "", err
	}
	log.Printf("Result: %s", c.String())
	return c.String(), nil
}
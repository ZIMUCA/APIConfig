package helpers

import (
	"bytes"
	"embed"
	"strings"
	"text/template"

	"github.com/adrg/frontmatter"
)

//go:embed config/*
var embeddedTemplates embed.FS

type MailMatter struct {
	Subject string `yaml:"subject"`
}

func GetStringFromEmbeddedTemplate(
	templatePath string,
	body interface{},
) (content string, matter MailMatter, err error) {

	tpl, err := template.ParseFS(embeddedTemplates, templatePath)
	if err != nil {
		return
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, body); err != nil {
		return
	}

	var parsed []byte
	parsed, err = frontmatter.Parse(strings.NewReader(buf.String()), &matter)
	if err != nil {
		return
	}

	content = string(parsed)
	return
}

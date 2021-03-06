package beego

import (
	"os"
	"path/filepath"
	"testing"
)

var header string = `{{define "header"}}
<h1>Hello, astaxie!</h1>
{{end}}`

var index string = `<!DOCTYPE html>
<html>
  <head>
    <title>beego welcome template</title>
  </head>
  <body>
{{template "block"}}
{{template "header"}}
{{template "blocks/block.tpl"}}
  </body>
</html>
`

var block string = `{{define "block"}}
<h1>Hello, blocks!</h1>
{{end}}`

func TestTemplate(t *testing.T) {
	dir := "_beeTmp"
	files := []string{
		"header.tpl",
		"index.tpl",
		"blocks/block.tpl",
	}
	if err := os.MkdirAll(dir, 0777); err != nil {
		t.Fatal(err)
	}
	for k, name := range files {
		os.MkdirAll(filepath.Dir(filepath.Join(dir, name)), 0777)
		if f, err := os.Create(filepath.Join(dir, name)); err != nil {
			t.Fatal(err)
		} else {
			if k == 0 {
				f.WriteString(header)
			} else if k == 1 {
				f.WriteString(index)
			} else if k == 2 {
				f.WriteString(block)
			}

			f.Close()
		}
	}
	if err := BuildTemplate(dir); err != nil {
		t.Fatal(err)
	}
	if len(BeeTemplates) != 3 {
		t.Fatalf("should be 3 but got %v", len(BeeTemplates))
	}
	if err := BeeTemplates["index.tpl"].ExecuteTemplate(os.Stdout, "index.tpl", nil); err != nil {
		t.Fatal(err)
	}
	for _, name := range files {
		os.RemoveAll(filepath.Join(dir, name))
	}
	os.RemoveAll(dir)
}

package examples

import (
	"text/template"
	"fmt"
	"os"
)

type DBInfo  struct {
	Dbname string
	Csvname string
	Tablename string
}

func LoadTemplate() {
	const template_text = `import -separator , ".import  {{.Dbname}}  {{.Csvname}} {{.Tablename}}"`
	tpl := template.Must(template.New("template").Parse(template_text))

	member := DBInfo{"HackGeo.db", "Counties.csv", "counties"}

	if err := tpl.Execute(os.Stdout, member); err != nil {
		fmt.Println(err)
	}

}


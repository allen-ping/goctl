package gen

import (
	"fmt"

	"github.com/allen-ping/goctl/v2/model/sql/template"
	"github.com/allen-ping/goctl/v2/util"
	"github.com/allen-ping/goctl/v2/util/pathx"
)

func genNew(table Table, withCache, postgreSql bool) (string, error) {
	text, err := pathx.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source()) + "`"
	}

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]any{
			"table":                 t,
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

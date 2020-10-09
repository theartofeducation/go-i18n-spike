package main

import (
	"strings"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// First parameter: Glob filepath pattern
	// Second variadic parameter: Optional language tags
	// The first language is the default/fallback
	app.I18n.Load("./locales/*/*", "en-US", "el-GR", "zh-CN")

	// Load all languages and set default/fallback
	// app.I18n.Load("./locales/*/*")
	// app.I18n.SetDefault("en-US")

	// Embed locales with go-bindata
	// go get github.com/go-bindata/go-bindata/...
	// go-bindata -o locales.go ./locales/...
	// app.I18n.LoadAssets(AssetNames, Asset, "en-US", "el-GR", "zh-CN")

	ctx.Tr("YouLate", 1)
	ctx.Tr("YouLate", 10)
	ctx.Tr("HouseCount", 2, "John")

	ctx.Tr("FreeDay", 3, 15)
	ctx.Tr("FreeDay", 1)
	ctx.Tr("FreeDay", 5)

	const (
		female = iota + 1
		nonbinary
	)

	// Outputs: Pat, Drew are awesome.
	ctx.Tr("VarTemplatePlural", iris.Map{
		"PluralCount": 5,
		"Names":       []string{"Pat", "Drew"},
		"InlineJoin": func(arr []string) {
			return strings.Join(arr, ", ")
		},
	})

	// Outputs: She is awesome!
	ctx.Tr("VarTemplatePlural", 1, female)

	// Outputs: They have 5 houses.
	ctx.Tr("VarTemplatePlural", 2, female, 5)

	// Outputs: Welcome John
	ctx.Tr("Welcome.Message", iris.Map{"Name": "John"})
}

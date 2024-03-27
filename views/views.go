package views

import (
	"html/template"
)

// FuncMap 用于在模板中添加数学运算的辅助函数
var FuncMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
}

var ColumnColors = []string{
	"#ffdddd", "#ddffdd", "#ddddff",
	"#ffddff", "#ddffff", "#ffffdd",
	"#d3d3d3", "#ffa07a", "#20b2aa",
}

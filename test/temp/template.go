package temp

import (
	"bytes"
	"fmt"
	"text/template"
)

func FormatUser() {
	// 模板定义
	tepl := `NAME: {{.Name}}, AGE: {{.Age}}`

	// 解析模板
	tmpl, _ := template.New("test").Parse(tepl)

	// 数据模型
	type User struct {
		Name string
		Age  uint
	}

	user := User{"Alice", 18}
	var result bytes.Buffer
	_ = tmpl.Execute(&result, user)

	fmt.Println(result.String())
}

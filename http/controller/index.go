package controller

import (
	"fmt"
	"github.com/wobusbz/wohuarou/global"
	"html/template"
	"net/http"
)

func (c *BaseController) Index(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(global.App.RootDir + "/template/index.html")
	if err != nil {
		fmt.Fprint(writer, "解析模板错误....")
	}
	t.Execute(writer, nil)

	fmt.Fprintf(writer, "test......")

}

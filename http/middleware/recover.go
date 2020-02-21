package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func Recover(f func(writer http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1024)
				runtime.Stack(buf, false)

				log.Printf("recover, error:%s\n%s\n", err, buf)
				fmt.Fprint(writer, "服务异常：", err)
			}
		}()
		f(writer, req)
	}
}

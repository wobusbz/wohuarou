package middleware

import (
	"github.com/wobusbz/wohuarou/util"
	"log"
	"net/http"
	"time"
)

func Logging(f func(writer http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		w := &statusWriter{ResponseWriter: writer}
		defer func() {
			elapse := time.Now().Sub(startTime)
			log.Printf(`[%s %s %s "%s" %d %d %s]`,
				util.RemoteIp(req),
				req.Method,
				req.RequestURI,
				req.UserAgent(),
				w.length,
				w.status,
				elapse)
		}()
		f(w, req)
	}
}

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *statusWriter) Write(b []byte) (int, error) {
	// 默認不呼叫 WriteHeader, 除非錯誤
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func (w *statusWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

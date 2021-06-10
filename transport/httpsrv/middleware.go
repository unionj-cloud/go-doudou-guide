package httpsrv

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		logrus.Infof(
			"%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func Rest(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		inner.ServeHTTP(w, r)
	})
}

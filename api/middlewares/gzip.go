// Created by nazarigonzalez on 15/10/17.

package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

//todo improve with pool sync to reuse objects

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		next.ServeHTTP(gzr, r)
	}
}

package httpUtil

import (
	"io"
	"net/http"
)

func write(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "xxxxxx")
}
func StartServer(srv *http.Server) error {
	http.HandleFunc("/test", write)
	return srv.ListenAndServe()
}

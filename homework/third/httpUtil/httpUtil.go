package httpUtil

import (
	"io"
	"net/http"
)

func write(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "xxxxxx")
}
func StartServer(srv *http.Server) error {
	http.HandleFunc("/test", write)
	return srv.ListenAndServe()
}

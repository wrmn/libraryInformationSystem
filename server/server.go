package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"librarySysfo/util"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Starting backend service
func Serve() {
	key = []byte(fmt.Sprintf("%x", md5.Sum([]byte(os.Getenv("SECRET_KEY")))))
	r := mux.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv(("SERVICE_PORT")))

	r.HandleFunc("/auth/login", login).Methods("POST")
	r.HandleFunc("/auth/refresh", refresh).Methods("POST")
	r.HandleFunc("/me", dashboard).Methods("GET")
	r.NotFoundHandler = notFoundHandler()
	r.MethodNotAllowedHandler = notAllowedHandler()

	util.InfoPrint(1, fmt.Sprintf("service at port %s", port))
	http.ListenAndServe(port, r)
}

func responseFormatter(r responseParam) {
	r.W.Header().Set("Content-Type", "application/json")
	r.W.WriteHeader(r.Status)
	r.W.Write(r.Body)
}

func respToByte(s string, m interface{}, c int) []byte {
	msg, _ := json.Marshal(response{Status: s, Message: m, Code: c})
	return msg
}

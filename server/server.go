package server

import (
	"net/http"
)

func LocalHost() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

package handlers

import (
	"fmt"
	"net/http"
)

func GetHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

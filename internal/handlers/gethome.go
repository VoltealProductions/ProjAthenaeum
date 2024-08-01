package handlers

import (
	"fmt"
	"net/http"
)

func GetHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Home Page!")
}

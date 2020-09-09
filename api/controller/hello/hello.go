package hello

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gos!")
}

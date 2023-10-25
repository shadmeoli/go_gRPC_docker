package main

import (
	"controller/index"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func entry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "")
}

func main() {
	fmt.Println("home")
}

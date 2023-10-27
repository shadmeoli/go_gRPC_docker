package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"net/http"
)

func entry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "")
}

func main() {
	fmt.Println("home")
}

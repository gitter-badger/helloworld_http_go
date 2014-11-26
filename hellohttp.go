package main

import (
 	"flag"
	"fmt"
	"net/http"
)

func main() {
  port := flag.Int("port", 8080, "a porta que você irá utilizar para acessar o serviço.")
  
  flag.Parse()
  
  http.HandleFunc("/", helloHandler)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Olá mundo!")
}

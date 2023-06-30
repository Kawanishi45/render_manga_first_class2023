package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/render_manga_api/controller"
	"github.com/render_manga_api/model/repository"
)

var (
	tr = repository.NewTitleRepository()
	tc = controller.NewTitleController(tr)

	pr = repository.NewPageRepository()
	pc = controller.NewPageController(pr)

	ro = controller.NewRouter(tc, pc)
)

func main() {
	var err error
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/home", ro.HandleTitlesRequest)
	http.HandleFunc("/viewer", ro.HandlePagesRequest)
	fmt.Println("ListenAndServe...")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

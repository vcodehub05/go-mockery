package main

import (
	"fmt"
	"log"
	"net/http"

	"go-mockery/api"
	"go-mockery/user"
	"go-mockery/user/repository"
)

func main() {

	repo := repository.NewRepo()
	svc := user.NewService(repo)
	r := api.Router(svc)

	fmt.Println("Starting server on the port 9000")

	log.Fatal(http.ListenAndServe("127.0.0.1:9000", r))

}

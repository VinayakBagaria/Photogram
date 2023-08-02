package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
	"github.com/VinayakBagaria/go-cat-pictures/api/routes"
	"github.com/VinayakBagaria/go-cat-pictures/config"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/gorilla/mux"
)

func main() {
	config.Init()

	repository := repository.NewPicturesRepository()
	service := service.NewPicturesService(repository)
	handlers := resthandlers.NewPicturesHandlers(service)
	routesList := routes.NewPicturesRoutes(handlers)

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, routesList)

	apiPort, err := strconv.Atoi(config.GetEnvString("server.port"))
	if err != nil {
		log.Fatalln("Unable to parse api port")
	}

	log.Printf("API service running on port: %d", apiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router))
}

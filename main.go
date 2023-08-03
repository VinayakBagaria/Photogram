package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VinayakBagaria/go-cat-pictures/api/resthandlers"
	"github.com/VinayakBagaria/go-cat-pictures/api/routes"
	"github.com/VinayakBagaria/go-cat-pictures/config"
	"github.com/VinayakBagaria/go-cat-pictures/db"
	"github.com/VinayakBagaria/go-cat-pictures/repository"
	"github.com/VinayakBagaria/go-cat-pictures/service"
	"github.com/VinayakBagaria/go-cat-pictures/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	dbConfig := db.NewConfiguration()
	dbHandler, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Panicln(err)
	}

	repository := repository.NewPicturesRepository(dbHandler)
	localStorage := storage.NewStorage(config.GetConfigValue("server.imagePath"))
	service := service.NewPicturesService(repository, localStorage)
	handler := resthandlers.NewPicturesHandler(service)
	routesList := routes.NewPicturesRoutes(handler)

	serverHandler := resthandlers.NewServerHandler()
	serverRoutesList := routes.NewServerRouteList(serverHandler)

	router := gin.Default()
	// Logger middleware will write the logs to gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/get-image", "./images")

	routes.Install(router, routesList)
	routes.Install(router, serverRoutesList)

	apiPort, err := strconv.Atoi(config.GetConfigValue("server.port"))
	if err != nil {
		log.Fatalln("Unable to parse api port")
	}

	log.Printf("API service running on port: %d", apiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router))
}

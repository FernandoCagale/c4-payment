package main

import (
	"fmt"
	"github.com/FernandoCagale/c4-payment/api/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	session, err := SetupMongoDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Erro to start MongoDB")
	}

	defer session.Close()

	app, err := SetupApplication(session)
	if err != nil {
		panic("Erro to start application")
	}

	go app.MakeEvents()

	router := app.MakeHandlers()

	router.Use(middleware.Header)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


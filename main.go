package main

import (
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

)

func main(){

	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString==""{
		log.Fatal("Port is not found")
	}
	
	log.Printf("Server starting on port %v",portString)
	router:= chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	v1Router:=chi.NewRouter()

	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/err",handleErr)
	  router.Mount("/v1",v1Router)

	server:=&http.Server{
		Handler:router,
		Addr:":"+portString,
	}
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}
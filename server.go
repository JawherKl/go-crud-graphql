package main

import (
	"log"
	"net/http"
	"os"

    "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/handlers"
	"github.com/jawherkl/go-crud-graphql/graph"
	"github.com/jawherkl/go-crud-graphql/graph/generated"
	"github.com/joho/godotenv"
	)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := godotenv.Load(".env");
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Connexion à la base de données
	Database := graph.Connect()

    	// Création du serveur GraphQL
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

	// Configuration des en-têtes CORS
    	corsHandler := handlers.CORS(
        	handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Autoriser les requêtes depuis ce domaine
        	handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    	)

    	// Configuration des routes
    	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    	http.Handle("/query", corsHandler(srv))

    	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    	log.Fatal(http.ListenAndServe(":"+port, nil))
}

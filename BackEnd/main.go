package main

import (
	"Backend/Database"
	create "Backend/Server"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

func main() {
	db, err := database.NewEntDatabase()
	s := create.NewServer(&db)
	if err != nil {
		fmt.Println("Failed to initialize database: ")
		return
	}
	fmt.Println("Database is ready")
	s.ApplyRoutes()
	if err := s.Router.Run(":" + os.Getenv("API_PORT")); err != nil {
		log.Fatalln(err)
	}
}

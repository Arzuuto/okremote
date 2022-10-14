package server

import (
	database "Backend/Database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Db     *database.Database
}

func NewServer(d *database.Database) Server {
	return Server{
		Router: gin.Default(),
		Db:     d,
	}
}

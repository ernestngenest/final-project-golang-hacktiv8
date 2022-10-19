package server

import (
	"fmt"
	"os"

	db "final_project_hacktiv8/connection" // import connection

	"github.com/gin-gonic/gin"
)

func Start() error {
	db, err := db.New()
	if err != nil {
		return err
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%s", port))
	return nil
}

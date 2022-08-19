package main

import (
	"github.com/aldotobing/go-gin-api-with-postgres/pkg/books"
	"github.com/aldotobing/go-gin-api-with-postgres/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	// register more routes here

	// r.GET("/books", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": "Aldo",
	// 		"bio":  "A Software Engineer",
	// 	})
	// })

	r.Run(port)
}

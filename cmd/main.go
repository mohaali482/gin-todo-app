package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/mohaali482/todo/pkg/config"
	"github.com/mohaali482/todo/pkg/database"
	"github.com/mohaali482/todo/pkg/routes"
)

func Init() (string, string) {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
		return "", ""
	}

	port := config.Port
	dbURL := config.DB_URL

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return port, dbURL
}

func main() {
	port, dbURL := Init()
	gin.ForceConsoleColor()

	r := gin.Default()
	d := database.InitDatabase(dbURL)

	routes.TodoRoutes(r, d)

	r.Run(port)
}

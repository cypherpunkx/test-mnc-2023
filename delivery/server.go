package delivery

import (
	"gorm-practice/delivery/routes"

	"github.com/gin-gonic/gin"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := routes.SetupRouter(app.engine); err != nil {
		panic("Application Error!")
	}
}

func Server() *application {
	return &application{
		engine: gin.Default(),
	}
}

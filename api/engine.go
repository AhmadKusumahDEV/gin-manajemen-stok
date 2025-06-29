package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {

}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

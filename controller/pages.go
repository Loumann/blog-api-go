package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) LoginPage(con *gin.Context) {
	con.HTML(http.StatusOK, "LoginTitle.html", nil)
}

func (c *Controller) OnPointWindowLocation(con *gin.Context) {
	con.HTML(http.StatusOK, "BlogTemplate.html", nil)

}

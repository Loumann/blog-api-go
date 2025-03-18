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

func (c *Controller) SignUpPage(con *gin.Context) {
	con.HTML(http.StatusOK, "Registration.html", nil)
}

func (c *Controller) MyPage(con *gin.Context) {
	con.HTML(http.StatusOK, "MyPage.html", nil)
}
func (c Controller) SearchPage(con *gin.Context) {
	con.HTML(http.StatusOK, "Search.html", nil)

}

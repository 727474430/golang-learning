package rest

import (
	"fifa_server/rest/teams"
	"github.com/gin-gonic/gin"
)

func ApiServerInit(r *gin.Engine) {
	v1 := r.Group("/v1/api")
	teamsRouter(v1)

	r.Run(":5000")
}

func teamsRouter(r *gin.RouterGroup) {
	r.GET("/teams", teams.TeamListHandler)
	r.GET("/teams/:teamId", teams.TeamHandler)
}

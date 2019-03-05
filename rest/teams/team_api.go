package teams

import (
	"fifa_server/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	ErrorTeam      = errors.New("Team does not exists!")
	ErrorTeamId    = errors.New("Team id is null!")
	ErrorTeamParam = errors.New("Team list query fail, because parameter not illegal!")
)

type ListTeamParam struct {
	Keyword string `from:"keyword"`
	Return  string `from:"return"`
}

func TeamListHandler(c *gin.Context) {
	var param ListTeamParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.AbortWithError(400, ErrorTeamParam)
		return
	}

	teams := make([]domain.TeamSerializer, 10)
	c.JSON(200, teams)
}

func TeamHandler(c *gin.Context) {
	teamId := c.Param("teamId")
	if teamId == "" {
		c.AbortWithError(400, ErrorTeamId)
		return
	}

	var team domain.Team
	c.JSON(200, team.Serializer())
}

package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Flag string `gorm:"type:varchar(128); not null; column:flag_address"`
	Name string `gorm:"type:varchar(64); not null; column:team_name"`
}

type TeamSerializer struct {
	ID         uint      `json:"id"`
	CreateDate time.Time `json:create_date`
	UpdateDate time.Time `json:update_date`
	Flag       string    `json:flag`
	Name       string    `json:name`
}

func (t *Team) Serializer() TeamSerializer {
	return TeamSerializer{
		ID:         t.ID,
		CreateDate: t.CreatedAt,
		UpdateDate: t.UpdatedAt,
		Flag:       t.Flag,
		Name:       t.Name,
	}
}

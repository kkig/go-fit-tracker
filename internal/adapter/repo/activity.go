package repo

import (
	"go-fit-tracker/internal/adapter/postgres"
	"go-fit-tracker/internal/adapter/postgres/models"
	"go-fit-tracker/internal/domain"
)

type Activity struct{}

func (a Activity) Create(activity domain.Activity) error {
	db := postgres.Connect()

	var user	models.User
	var exc		models.Excersize

	// Is user id valid?
	err := db.Where("id=?", activity.User.ID).Find(&user).Error
	if err != nil {
		return err
	}

	// Is exc id valid?
	err  = db.Where("id?=", activity.Excersize.ID).Find(&exc).Error
	if err != nil {
		return err
	}

	newData := models.Activity{
		ID:			activity.ID,
		UserID:		activity.User.ID,
		ExcID:		activity.Excersize.ID,
		Duration:	activity.Duration,
		Date:		activity.Date,		
	}

	err = db.Create(&newData).Error
	if err != nil {
		return err
	}

	return nil
}
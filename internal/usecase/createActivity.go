package usecase

import (
	"go-fit-tracker/internal/domain"
	"go-fit-tracker/internal/domain/repo"

	"github.com/google/uuid"
)

var (
	ActivityRepo	repo.IActivity
)

func CreateActivity(input domain.ActInput) domain.Activity {
	activity := ActivityRepo.Get()
	
	activity.ID				= uuid.New().String()
	activity.User.ID 		= input.UserID
	activity.Excersize.ID	= input.ExcID
	activity.Duration		= input.Duration
	activity.Date			= input.Date

	return activity
}

// type CreateActivityArgs struct {
// 	UserID			string
// 	Excersize		string
// 	Duration		int
// 	Date      		string
// 	UserRepo		repo.IUser
// 	ActivityRepo	repo.IActivity
// }

// func CreateActivity(args CreateActivityArgs) (domain.Activity, error) {
// 	user := args.UserRepo.GET(args.UserID)
	
// 	excType, err := valueobj.ConvertToExcType(args.Excersize)
// 	if err != nil {
// 		return domain.Activity{}, err
// 	}
// 	exc := valueobj.Excersize{ExcType: excType}

// 	activity := domain.Activity{
// 		ID:        	"123",
// 		User:      	user,
// 		Excersize:	exc,
// 		Duration:  	args.Duration,
// 		Date:      	args.Date,		
// 	}

// 	args.ActivityRepo.Save(activity)
// 	return activity, nil
// }
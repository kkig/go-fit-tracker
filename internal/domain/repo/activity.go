package repo

import "go-fit-tracker/internal/domain"

type IActivity interface {
	Get()	domain.Activity
}
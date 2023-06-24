package repo

import "go-fit-tracker/internal/domain"

type IActInput interface {
	Get()	domain.ActInput
}
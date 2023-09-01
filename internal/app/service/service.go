package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (*Service) DaysLeft() int64 {
	comparableMoment := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeLeft := time.Until(comparableMoment)
	return int64(timeLeft.Hours() / 24)
}

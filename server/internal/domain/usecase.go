package domain

import "log"

type Auth interface {
	Do()
}

type UseCase struct {
	auth Auth
}

func (uc UseCase) Auth() {
	uc.auth.Do()
}

func (uc UseCase) Mock() {
	log.Println("mock call usecase")
}

func New(au Auth) *UseCase {
	return &UseCase{
		auth: au,
	}
}

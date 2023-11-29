package app

import (
	"server/internal/domain"
	"server/internal/infrastructure/authorization"
	"server/internal/server/http"
)

type App struct {
	httpServer *http.Http
	uc         *domain.UseCase
	au         *authorization.Authorization
}

func (a *App) Run() {
	a.au = authorization.New()
	a.uc = domain.New(a.au)
	a.httpServer = http.New(a.uc)

	a.httpServer.Run()
}

func New() *App {
	return &App{}
}

package internal

type App struct {
	handler Handler
	store   Storage
}

func Initialize(key string) *App {
	return &App{
		handler: NewHandler(key),
		store:   NewStore(),
	}
}

func (a *App) Start(web, url, csrf string) {
	a.handler.Register(a.store, url, web, csrf)
}

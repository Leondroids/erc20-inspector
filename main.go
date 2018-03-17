package erc20_inspector

import (
	"github.com/Leondroids/erc20-inspector/app"
	"log"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/envion/golib"
	"github.com/Leondroids/gox"
	"net/http"
)

func main() {
	context, err := app.InitApp()

	if err != nil {
		panic(err)
	}

	log.Printf("Config: %+v\n", context.Config)

	gox.StartServer(router(context), context.Config.Port)
}

func router(context *app.Context) *mux.Router {
	commonHandlers := alice.New(golib.LoggingHandler)

	// main router
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/healthcheck", commonHandlers.ThenFunc(gox.HealthcheckHandler)).Methods("GET")

	// node
	tokenHandler := app.NewTokenHandler(context)
	nodeRouter := router.PathPrefix("/node").Subrouter().StrictSlash(true)
	nodeRouter.Handle("/status", commonHandlers.ThenFunc(tokenHandler.GetToken)).Methods("GET")

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	return router
}

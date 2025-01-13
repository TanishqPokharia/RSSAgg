package main

import (
	"RSSAggregator/handlers"
	"RSSAggregator/middlewares"
	"RSSAggregator/utils"
	"github.com/go-chi/chi/v5"
)

func handlerV1Router(apiConf *utils.ApiConfig) *chi.Mux {
	v1Router := chi.NewRouter()
	//v1Router.HandleFunc("/ready", handlerReadiness) // will respond to every ype of request
	v1Router.Route("/users", func(r chi.Router) {
		r.Get("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerGetUserByApiKey))
		r.Post("/", handlers.HandlerCreateUser(apiConf))
	})
	v1Router.Route("/feeds", func(r chi.Router) {
		r.Get("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerGetFeeds()))
		r.Post("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerCreateFeed()))
	})
	v1Router.Route("/feedFollows", func(r chi.Router) {
		r.Get("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerGetFeedFollows()))
		r.Post("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerCreateFeedFollow()))
		r.Delete("/{feedFollowId}", middlewares.AuthMiddleware(apiConf, handlers.HandlerDeleteFeedFollow()))
	})
	v1Router.Route("/posts", func(r chi.Router) {
		r.Get("/", middlewares.AuthMiddleware(apiConf, handlers.HandlerGetPosts))
	})
	v1Router.Get("/ready", handlers.HandlerReadiness) // explicitly mention request type
	return v1Router
}

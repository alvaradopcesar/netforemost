package http

import (
	"net/http"
	"netforemost/pkg/logger"
	notaHandler "netforemost/pkg/nota/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	status "netforemost/pkg/status/handler/http"
)

// routes function sets routes handlers.
func routes(log logger.Logger) http.Handler {
	r := chi.NewRouter()

	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
			"PATCH",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"Cache-Control",
		},
		AllowCredentials: true,
	})

	r.Use(co.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//r.Get("/swagger/*", httpSwagger.WrapHandler)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//})

	// Status handlers.
	sh := status.New(log)
	r.Get("/hello", sh.SayHelloHandler)
	r.Get("/health", sh.HealthCheckHandler)

	notaHttp := notaHandler.New(log)

	r.Post("/v1/nota", notaHttp.NoteCreateHandler)
	//r.Get("/v1/nota/{sort}", notaHttp.NoteGetAllHandler)
	r.Get("/v1/nota", notaHttp.NoteGetAllHandler)
	r.Put("/v1/nota", notaHttp.NoteUpdateHandler)

	// Add router V1
	//apiVersion1 := chi.NewRouter()
	//apiRouter.Get("/articles/{date}-{slug}", getArticle)
	//r.Mount("/v1/", getV1(dbClient, notification, log, tokenConf, cacheMidd, redisClient))
	//apiVersion1.Route("/v1/nota", func(r chi.Router) {
	//	apiVersion1.Post("/", notaHttp.NoteCreateHandler)
	//	apiVersion1.Get("/", notaHttp.NoteGetAllHandler) // GET /articles/123
	//	apiVersion1.Put("/", notaHttp.NoteUpdateHandler) // PUT /articles/123
	//})
	//r.Mount("/v1", apiVersion1)
	//r.

	return r
}

package main

import (
	"log"
	"log/slog"
	"net/http"

	"ordersystem/repository"
	"ordersystem/rest"
	"ordersystem/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "ordersystem/docs" // swagger docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title       Order System
// @description This system enables drink orders and should not be used for the forbidden Hungover Games.
func main() {
	// connect to S3
	s3, err := storage.CreateS3client()
	if err != nil {
		log.Fatalln("failed to connect to S3:", err)
	}

	// connect to DB
	db, err := repository.NewDatabaseHandler()
	if err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	// prepopulate data (drinks + initial S3 content)
	if err := repository.Prepopulate(db, s3); err != nil {
		log.Fatalln("failed to prepopulate database:", err)
	}

	// router
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// allow local CORS (frontend on localhost, API on orders.localhost via Traefik)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost",
			"http://localhost:80",
			"http://orders.localhost",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"Origin",
			"Cache-Control",
			"Expires",
			"Pragma",
		},
		ExposedHeaders:   []string{"Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300, // seconds
	}))

	// Simple health endpoint (useful for debugging / readiness checks)
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Menu Routes
	r.Get("/api/menu", rest.GetMenu(db))

	// Order Routes
	r.Get("/api/order/all", rest.GetOrders(db))
	r.Get("/api/order/totalled", rest.GetOrdersTotal(db))
	r.Get("/api/receipt/{orderId}", rest.GetReceiptFile(db, s3))
	r.Post("/api/order", rest.PostOrder(db, s3))

	// Swagger / OpenAPI UI
	// You can now use BOTH:
	//   http://orders.localhost/swagger/index.html
	//   http://orders.localhost/openapi/index.html
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/openapi/*", httpSwagger.WrapHandler)

	slog.Info("⚡⚡⚡ Order System is up and running on :3000 ⚡⚡⚡")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

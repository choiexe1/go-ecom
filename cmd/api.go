package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/choiexe1/go-ecom/internal/adapters/postgresql/sqlc"
	"github.com/choiexe1/go-ecom/internal/orders"
	ordersPostgres "github.com/choiexe1/go-ecom/internal/orders/postgres"
	"github.com/choiexe1/go-ecom/internal/products"
	productsPostgres "github.com/choiexe1/go-ecom/internal/products/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	queries := repo.New(app.db)

	productRepo := productsPostgres.NewRepository(queries)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	ordersRepo := ordersPostgres.NewRepository(queries, app.db)
	ordersService := orders.NewService(ordersRepo)
	ordersHandler := orders.NewHandler(ordersService)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		r.Route("/products", func(r chi.Router) {
			r.Get("/", productHandler.ListProduct)
			r.Post("/", productHandler.CreateProduct)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", productHandler.FindProductByID)
			})
		})

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", ordersHandler.PlaceOrder)
		})
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at address %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

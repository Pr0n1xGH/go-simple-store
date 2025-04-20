package main

import (
	"fmt"
	"go-start/pkg/database"
	"go-start/pkg/handler"
	"go-start/pkg/middleware"
	"go-start/pkg/repository"
	"go-start/pkg/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	vld := validator.New()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, vld)
	userHandler := handler.NewUserHandler(userService)

	cartItemRepo := repository.NewCartItemRepository(db)
	cartItemService := service.NewCartItemService(cartItemRepo)
	cartItemHandler := handler.NewCartItemHandler(cartItemService)

	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	authHandler := handler.NewAuthHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/api/user", userHandler.CreateUser)
		r.Get("/api/user", userHandler.GetUsers)
		r.Get("/api/user/{id}", userHandler.GetUser)
		r.Put("/api/user/{id}", userHandler.UpdateUser)
		r.Delete("/api/user/{id}", userHandler.DeleteUser)

		r.Post("/api/cart/create", cartHandler.CreateCart)
		r.Post("/api/cart/{cartID}/items", cartItemHandler.AddItem)
		r.Get("/api/cart/{cartID}/items", cartItemHandler.GetItems)
		r.Put("/api/cart/items/{itemID}", cartItemHandler.UpdateQuantity)
		r.Delete("/api/cart/items/{itemID}", cartItemHandler.DeleteItem)
	})

	r.Post("/api/login", authHandler.Login)
	r.Post("/api/refresh", authHandler.RefreshToken)

	fmt.Println("Сервер запущен на http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

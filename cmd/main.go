package main

import (
	"log"
	"net/http"
	"os"

	"reservation-system/internal/api/handler"
	"reservation-system/internal/api/middleware"
	"reservation-system/internal/infrastructure/db"
)

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	userHandler := handler.NewUserHandler()
	reservationHandler := handler.NewReservationHandler()
	authHandler := handler.NewAuthHandler()

	router := handler.NewRouter()

	router.POST("/api/auth/register", middleware.CORSMiddleware(authHandler.Register))
	router.POST("/api/auth/login", middleware.CORSMiddleware(authHandler.Login))
	router.POST("/api/auth/validate", middleware.CORSMiddleware(authHandler.ValidateToken))

	router.POST("/api/users", middleware.CORSMiddleware(userHandler.CreateUser))
	router.GET("/api/users", middleware.CORSMiddleware(userHandler.GetUser))
	router.POST("/api/users/login", middleware.CORSMiddleware(userHandler.Login))

	router.POST("/api/reservations", middleware.CORSMiddleware(middleware.AuthMiddleware(reservationHandler.CreateReservation)))
	router.GET("/api/reservations", middleware.CORSMiddleware(middleware.AuthMiddleware(reservationHandler.GetReservation)))
	router.GET("/api/reservations/user", middleware.CORSMiddleware(middleware.AuthMiddleware(reservationHandler.GetUserReservations)))
	router.POST("/api/reservations/confirm", middleware.CORSMiddleware(middleware.AuthMiddleware(reservationHandler.ConfirmReservation)))
	router.DELETE("/api/reservations", middleware.CORSMiddleware(middleware.AuthMiddleware(reservationHandler.CancelReservation)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

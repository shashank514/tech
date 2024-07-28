package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const version = "0.0.1"

// Dummy user data for authentication
var dummyUser = struct {
	Username string
	Password string
}{
	Username: "user",
	Password: "pass",
}

type config struct {
	port int
}

type application struct {
	config config
	logger *logrus.Logger
}

func main() {

	var cfg config

	// Try to read environment variable for port (given by railway). Otherwise use default
	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 4000
	}

	// Set the port to run the API on
	cfg.port = intPort

	// Create the logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Create the application
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	app.routes(r)

	// Create the server
	serverAddr := fmt.Sprintf(":%d", cfg.port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("server started", "addr", serverAddr)

	// Start the server
	if err := srv.ListenAndServe(); err != nil {
		logger.WithError(err).Error("server error")
		os.Exit(1)
	}
}

func (app *application) routes(r *gin.Engine) {
	r.POST("/login", app.handleLogin)
}

// handleLogin handles user login requests
func (app *application) handleLogin(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the JSON request body to the credentials struct
	if err := c.ShouldBindJSON(&credentials); err != nil {
		app.logger.WithError(err).Error("invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Mock authentication
	if credentials.Username == dummyUser.Username && credentials.Password == dummyUser.Password {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

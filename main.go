package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	"github.com/joho/godotenv"

	userHandler "github.com/todo-list/handler/users"
	"github.com/todo-list/middleware"
	userSvc "github.com/todo-list/service/users"
	userStore "github.com/todo-list/store/users"

	taskHandler "github.com/todo-list/handler/tasks"
	taskService "github.com/todo-list/service/tasks"
	taskStore "github.com/todo-list/store/tasks"
)

func init() {
	if os.Getenv("WORKING_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("error:%s", err)
		}
	}
}

func main() {
	var (
	// dbUser = os.Getenv("DB_USER")
	// dbPass = os.Getenv("DB_PASSWORD")
	// dbHost = os.Getenv("DB_HOST")
	// dbPort = os.Getenv("DB_PORT")
	// dbName = os.Getenv("DB_NAME")
	)
	// connString := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err.Error())
	}

	err = RunMigrations(db)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrations failed due to error %v", err.Error())
	}

	usrStore := userStore.New(db)
	usrService := userSvc.New(usrStore)
	usrHandler := userHandler.New(usrService)

	tskStore := taskStore.New(db)
	tskService := taskService.New(tskStore)
	tskHandler := taskHandler.New(tskService)

	r := gin.Default()

	// Register routes
	r.POST("/api/auth/register", usrHandler.Register)
	r.POST("/api/auth/login", usrHandler.Login)
	r.POST("/api/auth/forgot", usrHandler.ForgotPassword)
	r.POST("/api/auth/password/reset", usrHandler.ResetPassword)
	r.POST("api/auth/logout", usrHandler.Logout)

	// Protected routes (require authentication)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware()) // Apply AuthMiddleware to the protected routes
	{
		protected.GET("/users/current", usrHandler.GetCurrentUser)
		protected.PUT("/users/{id}", usrHandler.UpdateUserDetailsById)
		protected.POST("/tasks", tskHandler.CreateTask)
		protected.GET("/tasks", tskHandler.GetUserTasks)
		protected.GET("/tasks/:id", tskHandler.GetTaskById)
		protected.PUT("/tasks/:id", tskHandler.UpdateTaskById)
		protected.DELETE("/tasks/:id", tskHandler.DeleteTaskById)
		protected.PUT("/tasks/:id/mark", tskHandler.UpdateTaskCompletionStatus)
		protected.GET("/tasks/completed", tskHandler.GetUserCompletedTasks)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Start the server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server started on port ", port)

}

func RunMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"todo_medm", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	fmt.Println("migrations ran successfully")

	return nil
}

package main

import (
	"go-pagination/config"
	"go-pagination/controller"
	"go-pagination/repository"
	"go-pagination/routes"
	"go-pagination/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()

		siteRepository repository.SiteRepository = repository.NewSiteRepository(db)
		siteService    service.SiteService       = service.NewSiteService(siteRepository)
		siteController controller.SiteController = controller.NewSiteController(siteService)
	)

	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	routes.SiteRoutes(server, siteController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("error serving :", err)
	}
	// Graceful shutdown
	// 	go func() {
	// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 			log.Fatal("error serving :", err)
	// 		}
	// 	}()

	// 	quit := make(chan os.Signal, 1)

	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 	<-quit
	// 	log.Printf("[%v] - Shutting down server\n", time.Now())

	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	if err := srv.Shutdown(ctx); err != nil {
	// 		log.Fatal("error shutting down :", err)
	// 	}

	// 	<-ctx.Done()
	// 	log.Println("timeout, exiting")

}

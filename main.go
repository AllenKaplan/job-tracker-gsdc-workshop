package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"net/http"

	"github.com/AllenKaplan/jobTracker/db"
	"github.com/AllenKaplan/jobTracker/models"
)

type Server struct {
	Repo db.Repository
}

func NewServer(config db.Config) (Server, error) {
	repo, err := db.NewDB(config)
	if err != nil {
		return Server{}, err
	}
	return Server{Repo: repo}, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := db.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
	s, err := NewServer(config)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", pingHandler)
	r.GET("/", indexHandler)
	r.GET("/app/:id", s.getAppHandler)
	r.GET("/app", s.getAppsHandler)
	r.POST("/app", s.postAppHandler)
	r.PUT("/app", s.putAppHandler)
	r.DELETE("/app", s.deleteAppHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Job Tracker",
	})
}

func (s Server) getAppHandler(c *gin.Context) {
	var Id int
	c.Bind(&Id)
	res, err := s.Repo.GetApp(c.Request.Context(), Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) getAppsHandler(c *gin.Context) {
	res, err := s.Repo.GetApps(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) postAppHandler(c *gin.Context) {
	var App models.App
	c.Bind(&App)
	res, err := s.Repo.CreateApp(c.Request.Context(), App)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) putAppHandler(c *gin.Context) {
	var App models.App
	c.Bind(&App)
	res, err := s.Repo.UpdateApp(c.Request.Context(), App)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (s Server) deleteAppHandler(c *gin.Context) {
	var Id int
	c.Bind(&Id)
	res, err := s.Repo.DeleteApp(c.Request.Context(), Id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

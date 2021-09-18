package main

import (
	documents "chatoverchai/docs"
	"chatoverchai/storage"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	poetryManager := documents.DocumentManager{}
	blogManager := documents.DocumentManager{}
	previousManager := documents.DocumentManager{}
	upcomingManager := documents.DocumentManager{}
	photoManager := documents.DocumentManager{}
	videoManager := documents.DocumentManager{}

	storage.Load(&poetryManager, "poetry")
	storage.Load(&blogManager, "blog")
	storage.Load(&previousManager, "previous")
	storage.Load(&upcomingManager, "upcoming")
	storage.Load(&photoManager, "photo")
	storage.Load(&videoManager, "videos")

	fmt.Println(poetryManager)
	fmt.Println(blogManager)
	fmt.Println(previousManager)
	fmt.Println(upcomingManager)
	fmt.Println(photoManager)
	fmt.Println(videoManager)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Delims("[[", "]]")
	r.LoadHTMLGlob("*.html")

	r.Use(static.Serve("/js", static.LocalFile("./js", true)))
	r.Use(static.Serve("/images", static.LocalFile("./images", true)))
	r.Use(static.Serve("/videos", static.LocalFile("./videos", true)))
	r.Use(static.Serve("/TestVideo", static.LocalFile("./TestVideo", true)))

	r.GET("/poetry", func(c *gin.Context) {
		poetryManager = documents.DocumentManager{}
		storage.Load(&poetryManager, "poetry")
		c.HTML(http.StatusOK, "poetry.html", poetryManager)
	})

	r.GET("/blog", func(c *gin.Context) {
		blogManager := documents.DocumentManager{}
		storage.Load(&blogManager, "blog")
		c.HTML(http.StatusOK, "Blog2.html", blogManager)
	})

	r.GET("/upcoming", func(c *gin.Context) {
		upcomingManager := documents.DocumentManager{}
		storage.Load(&upcomingManager, "upcoming")
		c.HTML(http.StatusOK, "upcoming.html", upcomingManager)
	})

	r.GET("/previous", func(c *gin.Context) {
		previousManager := documents.DocumentManager{}
		storage.Load(&previousManager, "previous")
		c.HTML(http.StatusOK, "previous.html", previousManager)
	})

	r.GET("/photos", func(c *gin.Context) {
		photoManager := documents.DocumentManager{}
		storage.Load(&photoManager, "photo")
		c.HTML(http.StatusOK, "Photos.html", photoManager)
	})

	r.GET("/vids", func(c *gin.Context) {
		videoManager := documents.DocumentManager{}
		storage.Load(&videoManager, "videos")
		c.HTML(http.StatusOK, "vids.html", videoManager)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/home")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

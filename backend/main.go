package main

import (
	"net/http"

	"github.com/apichitlakorn/SA-PROJECT-DASHBOARD/backend/config"
	"github.com/apichitlakorn/SA-PROJECT-BB/tree/main/controller"
	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func main() {
    // Open connection to the database
    config.ConnectionDB()

    // Generate databases
    config.SetupDatabase()

    r := gin.Default()

    r.Use(CORSMiddleware())

    router := r.Group("")
    {
        // User 
        router.GET("/user/:id", controller.GetUserById) 
        // Reviews
        router.GET("/reviews", controller.ListReview) 
        router.POST("/reviews", controller.CreateReview) 
        router.GET("/reviews/course/:id", controller.GetReviewByCourseID) 
        router.GET("/reviews/filter", controller.GetFilteredReviews) 
        router.GET("/reviews/search", controller.SearchReviewsByKeyword) 
        router.GET("/course/:course_id/ratings", controller.GetRatingsAvgByCourseID) 
        router.POST("/reviews/like", controller.LikeReview)
        router.POST("/reviews/unlike", controller.UnlikeReview)
        router.GET("/reviews/:userID/:reviewID/like", controller.CheckUserLikeStatus)
    }

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
    })

    // Run the server
    r.Run("localhost:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

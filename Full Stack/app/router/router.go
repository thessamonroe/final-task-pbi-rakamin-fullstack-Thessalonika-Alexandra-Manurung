package router

import (
    "github.com/gin-gonic/gin"
    "fullStack/app/controllers"
    "fullStack/app/middlewares"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(middlewares.JWTAuthMiddleware())

    user := r.Group("/users")
    {
        user.POST("/register", controllers.RegisterUser)
        user.POST("/login", controllers.LoginUser)
        user.PUT("/:userId", controllers.UpdateUser)
        user.DELETE("/:userId", controllers.DeleteUser)
    }

    photos := r.Group("/photos")
    {
        photos.POST("/", controllers.CreatePhoto)
        photos.GET("/", controllers.GetPhotos)
        photos.PUT("/:photoId", controllers.UpdatePhoto)
        photos.DELETE("/:photoId", controllers.DeletePhoto)
    }

    return r
}

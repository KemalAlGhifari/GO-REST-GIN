package main
import "github.com/gin-gonic/gin"
import "GO-REST-GIN/models"
import "GO-REST-GIN/controllers/ProductController"

func main (){
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("api/products", ProductController.Index)
	r.GET("api/products/:id", ProductController.Show)
	r.POST("api/product", ProductController.Create)
	r.PUT("api/	/:id", ProductController.Update)
	r.DELETE("api/product", ProductController.Delete)

	r.Run()
}
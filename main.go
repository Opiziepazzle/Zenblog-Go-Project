package main

import (
"fmt"
  "os"
"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
    "github.com/opiziepazzle/zenblog/initializers"
     "github.com/opiziepazzle/zenblog/controllers"
    "github.com/opiziepazzle/zenblog/middleware"
 
    
)

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDatabase()
    initializers.SyncDB()
}





func main() {
    // Set Gin mode to release
    gin.SetMode(gin.ReleaseMode)
    
   
    
  
    // Initialize standard Go html template engine
    engine := html.New("./views/datas", ".html")

    
    
    // Create new router/server
    // r := gin.Default()

    
//set up app 
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    
    
    // Configure App
 app.Static("/", "./public")
 app.Use(middleware.RequireAuth)
   
 
 //routes
 Routes(app)









// Debug print statement
fmt.Println("Starting the application...")

    //start app
    app.Listen(":" + os.Getenv("PORT"))




    
}



func Routes(app *fiber.App){
	app.Get("/", controllers.IndexController)
	app.Get("/about", controllers.AboutController)
	app.Get("/category", controllers.CategoryController)
	app.Get("/contact", controllers.ContactController)
	app.Get("/searchresult", controllers.SearchController)
	app.Get("/singlepost", controllers.SingleController)
    app.Post("/comments", controllers.UsersCommentController)
    
}
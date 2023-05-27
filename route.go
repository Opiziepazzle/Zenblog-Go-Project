package main




import (
	"github.com/gofiber/fiber/v2"
	
	"github.com/opiziepazzle/zenblog/controllers"
	

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
    "github.com/opiziepazzle/zenblog/initializers"

)
	



func Routes(app *fiber.App){
	app.Get("/", controllers.IndexController)
	app.Get("/about", controllers.AboutController)
	app.Get("/category", controllers.CategoryController)
	app.Get("/contact", controllers.ContactController)
	app.Get("/searchresult", controllers.SearchController)
	app.Get("/singlepost", controllers.SingleController)
	app.Post("/comments", controllers.UsersCommentController)
}
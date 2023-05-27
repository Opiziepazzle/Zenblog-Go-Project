package controllers

import (
	"github.com/gofiber/fiber/v2"





	"github.com/opiziepazzle/zenblog/initializers"
	"github.com/opiziepazzle/zenblog/models"
	
)




func IndexController(c *fiber.Ctx) error{
	return c.Render("index", fiber.Map{

	})
}

func AboutController(c *fiber.Ctx) error{
	return c.Render("about", fiber.Map{

	})
}

func CategoryController(c *fiber.Ctx) error{
	return c.Render("category", fiber.Map{

	})
}

func ContactController(c *fiber.Ctx) error{
	return c.Render("contact", fiber.Map{

	})
}

func SearchController(c *fiber.Ctx) error{
	return c.Render("search-result", fiber.Map{

	})
}

func SingleController(c *fiber.Ctx) error{
	return c.Render("single-post", fiber.Map{

	})
}


//if testing on postman use this

// func UsersCommentController(c *fiber.Ctx) error{
// 	user := new(models.User)
// 	if err:= c.BodyParser(user); err !=nil {
// 		return c.Status(500).SendString(err.Error())
// 	}
// initializers.DB.Save(&user)
// return c.JSON(&user)

//  // Redirect to the root route
// //  return c.Redirect("/")
// }



// webApp use this
func UsersCommentController(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initializers.DB.Save(&user)
	return c.SendStatus(fiber.StatusOK) // Return a success status without any content
	
}









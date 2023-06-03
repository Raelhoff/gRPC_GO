package routes

import (
    "fiber-mongo-api/controllers" //add this
    "github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    //All routes related to users comes here
    app.Post("/user", controllers.CreateUser) //add this
    app.Get("/user/:userId", controllers.GetAUser) //add this
    app.Put("/user/:userId", controllers.EditAUser) //add this
    app.Delete("/user/:userId", controllers.DeleteAUser) //add this
 
    app.Get("/users", controllers.GetAllUsers) //add this
}

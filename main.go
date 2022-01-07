package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"simple.crm.system/database"
	"simple.crm.system/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/lead", lead.GetLeads)
	app.Get("api/lead/:id", lead.GetLead)
	app.Post("api/lead", lead.NewLead)
	app.Delete("api/lead/:id", lead.DeleteLead)
}

func seedDatabase(db *gorm.DB) {
	db.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrations completed successfully")
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	database.DBConn = database.Init()
	seedDatabase(database.DBConn)

	app.Listen(3000)
	defer database.Close()
}

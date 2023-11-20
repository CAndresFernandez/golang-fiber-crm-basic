package main

import (
	"fmt"

	"github.com/CAndresFernandez/go-fiber-crm-basic/database"
	"github.com/CAndresFernandez/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	// get all leads
	app.Get("/api/v1/lead", lead.GetLeads)
	// get one lead by id
	app.Get("/api/v1/lead/:id", lead.GetLead)
	// create new lead
	app.Post("/api/v1/lead", lead.NewLead)
	// delete a lead by id
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
}
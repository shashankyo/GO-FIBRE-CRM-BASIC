package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shashank/Go-fiber-crm-basic/database"
	"github.com/shashank/Go-fiber-crm-basic/lead"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}

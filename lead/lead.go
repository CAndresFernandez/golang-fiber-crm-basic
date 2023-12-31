package lead

import (
	"github.com/CAndresFernandez/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct{
	gorm.Model
	Name string			`json:"name"`
	Company string		`json:"company"`
	Email string		`json:"email"`
	Phone int			`json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	// parse the data in the body, send error if necessary or create
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	// define the variable lead of type Lead
	var lead Lead
	// find the first instance of Lead that matches the id
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found which matches ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
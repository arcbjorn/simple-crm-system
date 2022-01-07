package lead

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"simple.crm.system/database"
)

type Lead struct {
	gorm.Model
	Name    string `json: "name"`
	Company string `json: "company"`
	Email   string `json: "email"`
	Phone   string `json: "phone"`
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := ctx.BodyParser(lead); err != nil {
		ctx.Status(503).Send(err)
		return
	}

	db.Create(&lead)
	ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		ctx.Status(500).Send("No lead found with this ID")
		return
	}
	db.Delete(&lead)
	ctx.Send("Lead successfully deleted")
}

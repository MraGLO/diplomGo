package http

import (
	"diplomGo/pkg/model"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) CreatePricing(c *fiber.Ctx) error {

	var pricingTable model.PricingTable
	var dataFromPricingTable model.DataFromPricingTable
	var teacherData model.TeacherData

	id, err := strconv.Atoi(c.Params("tableFileID"))
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}
	if id <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или равно 0"})
	}

	tableFile, err := h.services.GetTableFilesByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if tableFile.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	pricingTable.Name = tableFile.Name
	pricingTable.Date = time.Now()

	pricingTable.ID, err = h.services.AddPricingTable(pricingTable)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	dataFromPricingTable.PricingTable = pricingTable

	teachers, err := h.services.GetAllTeachersFromPriceView(tableFile.ID)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	for _, teacher := range teachers {
		recordForPricing, err := h.services.GetAllRecordForPricingFromPriceView(tableFile.ID, teacher.ID)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
		}

		teacherData.Teacher = teacher
		teacherData.RecordForPricing = recordForPricing
		dataFromPricingTable.TeacherData = append(dataFromPricingTable.TeacherData, teacherData)
	}

	c.Status(201)
	return c.JSON(dataFromPricingTable)
}

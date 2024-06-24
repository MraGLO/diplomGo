package http

import (
	"log"

	"encoding/json"

	"diplomGo/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetTeacherPricingRecords(c *fiber.Ctx) error {
	var pricingTeacherStruct model.GetPricingTeacherStruct

	err := json.Unmarshal(c.Body(), &pricingTeacherStruct)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	teacherPricingRecord, err := h.services.GetAllPricingRecordFromPricingTAbleViewByData(pricingTeacherStruct)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	c.Status(200)
	return c.JSON(teacherPricingRecord)
}

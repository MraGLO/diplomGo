package http

import (
	"diplomGo/pkg/model"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllRecordStudyPlans(c *fiber.Ctx) error {
	recordStudyPlans, err := h.services.GetAllRecordStudyPlans()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(recordStudyPlans)
}

func (h *Handlers) GetRecordStudyPlanById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("recordStudyPlanID"))
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

	recordStudyPlan, err := h.services.GetRecordStudyPlanByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if recordStudyPlan.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(recordStudyPlan)
}

func (h *Handlers) AddRecordStudyPlan(c *fiber.Ctx) error {
	var recordStudyPlan model.RecordStudyPlan

	err := json.Unmarshal(c.Body(), &recordStudyPlan)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.AddRecordStudyPlan(&recordStudyPlan)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

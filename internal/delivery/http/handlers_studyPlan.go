package http

import (
	"diplomGo/pkg/model"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllStudyPlans(c *fiber.Ctx) error {
	studyPlans, err := h.services.GetAllStudyPlans()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(studyPlans)
}

func (h *Handlers) GetStudyPlanById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("studyPlanID"))
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

	studyPlan, err := h.services.GetStudyPlanByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if studyPlan.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(studyPlan)
}

func (h *Handlers) AddStudyPlan(c *fiber.Ctx) error {
	var studyPlan model.StudyPlan

	err := json.Unmarshal(c.Body(), &studyPlan)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.AddStudyPlan(&studyPlan)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

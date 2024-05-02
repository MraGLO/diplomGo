package http

import (
	"diplomGo/pkg/model"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllLoad(c *fiber.Ctx) error {
	loads, err := h.services.GetAllLoads()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(loads)
}

func (h *Handlers) GetLoadById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("loadID"))
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

	load, err := h.services.GetLoadByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if load.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(load)
}

func (h *Handlers) AddLoad(c *fiber.Ctx) error {
	var load model.Load

	err := json.Unmarshal(c.Body(), &load)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.AddLoad(&load)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

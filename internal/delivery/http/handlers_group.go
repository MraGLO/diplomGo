package http

import (
	"diplomGo/pkg/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllGroups(c *fiber.Ctx) error {
	groups, err := h.services.GetAllGroups()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(groups)
}

func (h *Handlers) GetGroupsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("groupID"))
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

	group, err := h.services.GetGroupByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if group.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(group)
}

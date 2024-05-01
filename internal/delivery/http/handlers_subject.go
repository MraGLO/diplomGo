package http

import (
	"fmt"
	"log"
	"strconv"

	"encoding/json"

	"diplomGo/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllSubjects(c *fiber.Ctx) error {
	subjects, err := h.services.GetAllSubject()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(subjects)
}

func (h *Handlers) GetSubjectById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("subjectID"))
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

	subject, err := h.services.GetSubjectByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if subject.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(subject)
}

func (h *Handlers) AddSubject(c *fiber.Ctx) error {
	var subject model.Subject

	err := json.Unmarshal(c.Body(), &subject)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateSubjectData(subject)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddSubject(subject)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("SubjectID")
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}

	if id <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или рано 0"})
	}

	var subject model.Subject

	err = json.Unmarshal(c.Body(), &subject)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateSubjectData(subject)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateSubject(id, subject)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func isValidateSubjectData(Subject model.Subject) (str string) {
	str, b := isValidString(Subject.SubjectName)
	if Subject.SubjectName == "" || b {
		str = fmt.Sprintf("названии: %s", str)
		log.Println(str)
	}
	return
}

func (h *Handlers) DeleteSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("SubjectID")
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

	found, err := h.services.DeleteSubject(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if !found {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	c.Status(200)
	return c.JSON(model.Error{Data: "Успешно удалено"})
}

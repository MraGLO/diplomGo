package http

import (
	"encoding/json"
	"log"
	"strconv"

	"diplomGo/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllTeacherSubject(c *fiber.Ctx) error {
	teacherSubject, err := h.services.GetAllTeacherSubject()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(teacherSubject)
}

func (h *Handlers) GetTeacherSubjectById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("teacherSubjectID"))
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

	teacherSubject, err := h.services.GetTeacherSubjectByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if teacherSubject.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(teacherSubject)
}

func (h *Handlers) AddTeacherSubjects(c *fiber.Ctx) error {
	var teacherSubject model.TeacherSubject

	err := json.Unmarshal(c.Body(), &teacherSubject)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.AddTeacherSubject(teacherSubject)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateTeacherSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("teacherSubjectID")
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

	var teacherSubject model.TeacherSubject

	err = json.Unmarshal(c.Body(), &teacherSubject)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.UpdateTeacherSubject(id, teacherSubject)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func (h *Handlers) DeleteTeacherSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("teacherSubjectID")
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

	found, err := h.services.DeleteTeacherSubject(id)
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

func (h *Handlers) GetTeacherSubjectViewById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("teacherID"))
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

	subjects, err := h.services.GetASubjectsFromTeacherSubjectViewByTeacherID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	return c.JSON(subjects)
}

func (h *Handlers) DeleteTeacherSubjectByData(c *fiber.Ctx) error {
	var teacherSubject model.TeacherSubject

	err := json.Unmarshal(c.Body(), &teacherSubject)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	found, err := h.services.DeleteTeacherSubjectByData(teacherSubject)
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

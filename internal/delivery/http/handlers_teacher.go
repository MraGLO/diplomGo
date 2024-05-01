package http

import (
	"fmt"
	"log"
	"strconv"

	"encoding/json"

	"diplomGo/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAllTeachers(c *fiber.Ctx) error {
	teachers, err := h.services.GetAllTeachers()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(teachers)
}

func (h *Handlers) GetTeacherById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("TeacherID"))
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

	teacher, err := h.services.GetTeacherByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if teacher.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(teacher)
}

func (h *Handlers) AddTeacher(c *fiber.Ctx) error {
	var teacher model.Teacher

	err := json.Unmarshal(c.Body(), &teacher)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateTeacherData(teacher)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddTeacher(teacher)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateTeacher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("TeacherID")
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

	var teacher model.Teacher

	err = json.Unmarshal(c.Body(), &teacher)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateTeacherData(teacher)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateTeacher(id, teacher)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func isValidateTeacherData(Teacher model.Teacher) (str string) {
	str, b := isValidString(Teacher.Name)
	if Teacher.Name == "" || b {
		str = fmt.Sprintf("названии: %s", str)
		log.Println(str)
	}

	str, b = isValidString(Teacher.Surname)
	if Teacher.Surname == "" || b {
		str = fmt.Sprintf("названии: %s", str)
		log.Println(str)
	}

	str, b = isValidString(Teacher.Patronymic)
	if b {
		str = fmt.Sprintf("названии: %s", str)
		log.Println(str)
	}

	return
}

func (h *Handlers) DeleteTeacher(c *fiber.Ctx) error {
	id, err := c.ParamsInt("TeacherID")
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

	found, err := h.services.DeleteTeacher(id)
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

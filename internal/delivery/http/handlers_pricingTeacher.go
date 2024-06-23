package http

import (
	"encoding/json"
	"log"
	"strconv"

	"diplomGo/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetTeacherByPricingId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("pricingID"))
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

	teachers, err := h.services.GetAllTeachersFromPricingTableView(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	c.Status(200)
	return c.JSON(teachers)
}

func (h *Handlers) AddPricingTeacher(c *fiber.Ctx) error {
	var pricingTeacherStruct model.GetPricingTeacherStruct
	var pricingRecord model.PricingRecord
	var teacherPricingRecord model.TeacherPricingRecord

	err := json.Unmarshal(c.Body(), &pricingTeacherStruct)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	firstSubject, err := h.services.GetFirstSubject()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	firstGroup, err := h.services.GetFirstGroup()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	pricingRecord.GroupID = firstGroup.ID
	pricingRecord.SubjectID = firstSubject.ID
	pricingRecord.PricingTableID = pricingTeacherStruct.PricingID

	pricingRecord.ID, err = h.services.AddPricingRecord(pricingRecord)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	teacherPricingRecord.TeacherID = pricingTeacherStruct.TeacherID
	teacherPricingRecord.PricingRecordID = pricingRecord.ID

	err = h.services.AddTeacherPricingRecord(teacherPricingRecord)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(200)
	return c.JSON("Успешно добавлено")
}

func (h *Handlers) DeleteTeacherPricingByTeacherId(c *fiber.Ctx) error {
	var pricingTeacherStruct model.GetPricingTeacherStruct

	err := json.Unmarshal(c.Body(), &pricingTeacherStruct)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	prcingRecords, err := h.services.GetAllTeacherPricingRecordFromPricingTAbleViewByData(pricingTeacherStruct)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	for _, pricingRecord := range prcingRecords {
		found, err := h.services.DeletePricingRecord(pricingRecord.PricingRecordID)
		if !found {
			log.Println(err)
			c.Status(404)
			return c.JSON(model.Error{Data: "Данных по данному id не существует"})
		}
	}
	c.Status(200)
	return c.JSON("Успешно удалено")
}

package http

import (
	"diplomGo/pkg/model"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) CreatePricing(c *fiber.Ctx) error {

	var pricingTable model.PricingTable
	var pricingRecord model.PricingRecord
	var teacherPricingRecord model.TeacherPricingRecord
	var tableFileID model.TableFile

	err := json.Unmarshal(c.Body(), &tableFileID)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	if tableFileID.ID <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или равно 0"})
	}

	tableFile, err := h.services.GetTableFilesByID(tableFileID.ID)
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

	teachers, err := h.services.GetAllTeachersFromTableFileView(tableFile.ID)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	for _, teacher := range teachers {
		recordForTableFileView, err := h.services.GetAllRecordForPricingFromTableFileView(tableFile.ID, teacher.ID)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
		}
		for _, record := range recordForTableFileView {
			pricingRecord.PricingTableID = pricingTable.ID
			pricingRecord.GroupID = record.GroupID
			pricingRecord.SubjectID = record.SubjectID
			pricingRecord.HoursFirstSemester = record.TimeSemesterOwn
			pricingRecord.HoursSecondSemester = record.TimeSemesterTwo
			tso, err := parseTime(record.TimeSemesterOwn)
			if err != nil {
				log.Println(err)
				c.Status(500)
				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
			}
			tst, err := parseTime(record.TimeSemesterTwo)
			if err != nil {
				log.Println(err)
				c.Status(500)
				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
			}
			pricingRecord.Theory = tso.Theory + tst.Theory
			pricingRecord.Practice = tso.Practice + tst.Practice
			pricingRecord.Total = tso.Total + tst.Total
			pricingRecord.FirstHalfYear = tso.Total
			pricingRecord.SecondHalfYear = tst.Total

			pricingRecord.ID, err = h.services.AddPricingRecord(pricingRecord)
			if err != nil {
				log.Println(err)
				c.Status(500)
				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
			}
			teacherPricingRecord.TeacherID = teacher.ID
			teacherPricingRecord.PricingRecordID = pricingRecord.ID

			err = h.services.AddTeacherPricingRecord(teacherPricingRecord)
			if err != nil {
				log.Println(err)
				c.Status(500)
				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
			}
		}

	}

	c.Status(201)
	return c.JSON("Успешно добавлено")
}

func parseTime(s string) (*model.TimeStruct, error) {
	var ts model.TimeStruct

	// Проверяем, пустая ли строка
	if s == "" {
		return &ts, nil
	}

	// Регулярное выражение для поиска чисел в строке
	re := regexp.MustCompile(`(\d+)`)

	// Ищем числа в строке
	matches := re.FindAllString(s, -1)

	// Если чисел больше 3-х, то ошибка
	if len(matches) > 3 {
		return nil, fmt.Errorf("invalid string format: %s", s)
	}

	// Парсим числа
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int: %s", match)
		}

		switch i {
		case 0:
			ts.Total = num
		case 1:
			ts.Theory = num
		case 2:
			ts.Practice = num
		}
	}

	return &ts, nil
}
func (h *Handlers) GetAllPricing(c *fiber.Ctx) error {
	pricing, err := h.services.GetAllPricingTable()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(pricing)
}

func (h *Handlers) UpdatePricing(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pricingTableID")
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

	var pricingTable model.PricingTable

	err = json.Unmarshal(c.Body(), &pricingTable)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	err = h.services.UpdatePricingTable(id, pricingTable)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func (h *Handlers) DeletePricing(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pricingTableID")
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

	found, err := h.services.DeletePricingTable(id)
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

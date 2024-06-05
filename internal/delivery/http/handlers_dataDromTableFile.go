package http

import (
	"diplomGo/pkg/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetDataFromTableFileById(c *fiber.Ctx) error {
	var dataFromTableFile model.DataFromTableFile
	var dataFromSheet model.DataFromSheet
	var dataFromSheets []model.DataFromSheet
	var records []model.Record

	id, err := strconv.Atoi(c.Params("tableFileID"))
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

	tableFile, err := h.services.GetTableFilesByID(id)
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

	dataFromTableFile.TableFile = tableFile

	tableFileSheets, err := h.services.GetAllTableFileSheetsByTableFileID(tableFile.ID)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	for _, tableFileSheet := range tableFileSheets {
		sheet, err := h.services.GetSheetByID(tableFileSheet.SheetID)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
		}
		if sheet.ID == 0 {
			log.Println(err)
			c.Status(404)
			return c.JSON(model.Error{Data: "Данных по данному id не существует"})
		}

		sheetRecords, err := h.services.GetAllSheetRecordsBySheetID(sheet.ID)
		if err != nil {
			log.Println(err)
			c.Status(500)
			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
		}

		for _, sheetRecord := range sheetRecords {
			record, err := h.services.GetRecordByID(sheetRecord.RecordID)
			if err != nil {
				log.Println(err)
				c.Status(500)
				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
			}
			if record.ID == 0 {
				log.Println(err)
				c.Status(404)
				return c.JSON(model.Error{Data: "Данных по данному id не существует"})
			}
			records = append(records, record)
		}
		dataFromSheet.Sheet = sheet
		dataFromSheet.Records = records
		dataFromSheets = append(dataFromSheets, dataFromSheet)
		records = nil
	}
	dataFromTableFile.DataFromSheets = dataFromSheets

	return c.JSON(dataFromTableFile)
}

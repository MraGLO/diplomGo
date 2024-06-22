package http

import (
	"diplomGo/pkg/model"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func (h *Handlers) AddFile(c *fiber.Ctx) error {

	var objTableFile model.TableFile
	var objRecord model.Record
	var objGroupRecords model.GroupRecords

	file, err := c.FormFile("excel")
	if err != nil {
		log.Println(err)
		c.Status(400)
	}
	openFile, err := file.Open()
	if err != nil {
		log.Println(err)
		c.Status(400)
	}
	defer openFile.Close()

	f, err := excelize.OpenReader(openFile)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	objTableFile.Name = file.Filename
	objTableFile.Date = time.Now()
	idTableFile, err := h.services.AddTableFile(objTableFile)
	if err != nil {
		log.Println(err)
		return err
	}

	var sheets = f.GetSheetList()
	for _, sheet := range sheets {

		groupID, err := h.services.GetIDByGroupName(sheet)
		if err != nil {
			log.Println(err)
			return err
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Println(err)
			return err
		}
		for i, row := range rows {
			if i != 0 {
				objRecord.SubjectID, err = h.services.GetIDBySubject(row[1])
				if err != nil {
					log.Println(err)
					return err
				}
				objRecord.TeacherID, err = h.services.GetIDByTeacherSurname(row[6])
				if err != nil {
					log.Println(err)
					return err
				}
				objRecord.TimeSemesterOwn = row[2]
				objRecord.TimeSemesterTwo = row[3]
				objRecord.TableFileID = idTableFile

				recordID, err := h.services.AddRecord(objRecord)
				if err != nil {
					log.Println(err)
					return err
				}

				objGroupRecords.GroupID = groupID
				objGroupRecords.RecordID = recordID
				err = h.services.AddGroupRecords(objGroupRecords)
				if err != nil {
					log.Println(err)
					return err
				}

			}

		}
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) GetAllTableFile(c *fiber.Ctx) error {
	tableFiles, err := h.services.GetAllTableFiles()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})

	}
	return c.JSON(tableFiles)
}

func (h *Handlers) GetTableFileById(c *fiber.Ctx) error {
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

	tableFiles, err := h.services.GetTableFilesByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if tableFiles.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(tableFiles)
}

func (h *Handlers) DeleteTableFile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("TableFileID")
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

	found, err := h.services.DeleteTableFile(id)
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

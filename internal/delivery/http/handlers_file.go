package http

import (
	"diplomGo/pkg/model"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func (h *Handlers) AddFile(c *fiber.Ctx) error {

	var objTableFile model.TableFile
	var objSheet model.Sheet
	var objRecord model.Record
	var objTableFileSheet model.TableFileSheet
	var objSheetRecords model.SheetRecords

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

		objSheet.GroupID, err = h.services.GetIDByGroupName(sheet)
		if err != nil {
			log.Println(err)
			return err
		}
		idSheet, err := h.services.AddSheet(objSheet)
		if err != nil {
			log.Println(err)
			return err
		}

		objTableFileSheet.TableFileID = idTableFile
		objTableFileSheet.SheetID = idSheet

		err = h.services.AddTableFileSheets(objTableFileSheet)
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

				idRecord, err := h.services.AddRecord(objRecord)
				if err != nil {
					log.Println(err)
					return err
				}

				objSheetRecords.SheetID = idSheet
				objSheetRecords.RecordID = idRecord
				err = h.services.AddSheetRecords(objSheetRecords)
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

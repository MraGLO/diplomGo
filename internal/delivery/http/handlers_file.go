package http

import (
	"diplomGo/pkg/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func (h *Handlers) AddFile(c *fiber.Ctx) error {
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
	var sheets = f.GetSheetList()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Println(err)
			return err
		}
		for i, row := range rows {
			if i != 0 {
				for _, colCell := range row {
					log.Print(colCell, "\t")
				}
				log.Println()
			}

		}
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

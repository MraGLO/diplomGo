package http

// import (
// 	"diplomGo/pkg/model"
// 	"log"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )
//
// func (h *Handlers) GetDataFromTableFileById(c *fiber.Ctx) error {
// 	var dataFromTableFile model.DataFromTableFile
// 	var groupData model.GroupData
// 	var arrayGroupData []model.GroupData
// 	var records []model.Record

// 	id, err := strconv.Atoi(c.Params("tableFileID"))
// 	if err != nil {
// 		log.Println(err)
// 		c.Status(400)
// 		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
// 	}

// 	if id <= 0 {
// 		log.Println("id <=0")
// 		c.Status(400)
// 		return c.JSON(model.Error{Data: "id не может быть меньше или равно 0"})
// 	}

// 	tableFile, err := h.services.GetTableFilesByID(id)
// 	if err != nil {
// 		log.Println(err)
// 		c.Status(500)
// 		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
// 	}

// 	if tableFile.ID == 0 {
// 		log.Println(err)
// 		c.Status(404)
// 		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
// 	}

// 	dataFromTableFile.TableFile = tableFile

// 	tableFileGroups, err := h.services.GetAllTableFileGroupsByTableFileID(tableFile.ID)
// 	if err != nil {
// 		log.Println(err)
// 		c.Status(500)
// 		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
// 	}

// 	for _, tableFileGroup := range tableFileGroups {
// 		group, err := h.services.GetGroupByID(tableFileGroup.GroupID)
// 		if err != nil {
// 			log.Println(err)
// 			c.Status(500)
// 			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
// 		}
// 		if group.ID == 0 {
// 			log.Println(err)
// 			c.Status(404)
// 			return c.JSON(model.Error{Data: "Данных по данному id не существует"})
// 		}

// 		groupRecords, err := h.services.GetAllGroupRecordsByGroupID(group.ID)
// 		if err != nil {
// 			log.Println(err)
// 			c.Status(500)
// 			return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
// 		}

// 		for _, groupRecord := range groupRecords {
// 			record, err := h.services.GetRecordByID(groupRecord.RecordID)
// 			if err != nil {
// 				log.Println(err)
// 				c.Status(500)
// 				return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
// 			}
// 			if record.ID == 0 {
// 				log.Println(err)
// 				c.Status(404)
// 				return c.JSON(model.Error{Data: "Данных по данному id не существует"})
// 			}
// 			records = append(records, record)
// 		}
// 		groupData.Group = group
// 		groupData.Records = records
// 		arrayGroupData = append(arrayGroupData, groupData)
// 		records = nil
// 	}
// 	dataFromTableFile.GroupData = arrayGroupData

// 	return c.JSON(dataFromTableFile)
// }

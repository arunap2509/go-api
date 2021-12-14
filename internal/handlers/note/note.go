package noteHandler

import (
	"fmt"
	"net/http"

	"github.com/arunap2509/notes-api/database"
	"github.com/arunap2509/notes-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GetNotes(ctx *fiber.Ctx) error {
	db := database.DB

	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return ctx.JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": "no notes found",
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{"status": http.StatusOK, "message": "notes found", "data": notes})
}

func CreateNote(ctx *fiber.Ctx) error {

	db := database.DB

	note := new(model.Note)

	err := ctx.BodyParser(note)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "couldnt save note",
			"data":    err,
		})
	}

	fmt.Println(note)
	err = db.Create(note).Error

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "error while saving data",
			"data":    err,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "note saved successfully",
		"data":    note,
	})
}

func UpdateNote(ctx *fiber.Ctx) error {

	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"subTitle"`
		Text     string `json:"text"`
	}

	db := database.DB

	var updateNoteData updateNote
	var note model.Note

	err := ctx.BodyParser(&updateNoteData)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "error while parsing data",
			"data":    err,
		})
	}

	id := ctx.Params("id")

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "no data found",
			"data":    nil,
		})
	}

	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return ctx.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "data updated successfully",
		"data":    note,
	})
}

func GetNoteById(ctx *fiber.Ctx) error {

	db := database.DB

	id := ctx.Params("id")

	var note model.Note

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "no data found",
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{"status": fiber.StatusOK, "message": "data for" + id, "data": note})
}

func DeleteNote(ctx *fiber.Ctx) error {

	db := database.DB

	id := ctx.Params("id")

	var note model.Note

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "no data found",
			"data":    nil,
		})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "error while deleting data",
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{"status": fiber.StatusOK, "message": "data deleted successfully", "data": nil})
}

package noteRoutes

import (
	noteHandler "github.com/arunap2509/notes-api/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNotesRoute(router fiber.Router) {

	note := router.Group("notes")

	// create a note
	note.Post("/", noteHandler.CreateNote)

	// read all notes
	note.Get("/", noteHandler.GetNotes)

	// get a note by its ID
	note.Get("/:id", noteHandler.GetNoteById)

	// update a note
	note.Put("/:id", noteHandler.UpdateNote)

	// delete a note
	note.Delete("/:id", noteHandler.DeleteNote)
}

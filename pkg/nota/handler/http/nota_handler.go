package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"netforemost/pkg/logger"
	notaService "netforemost/pkg/nota/service"
	"netforemost/pkg/response"
)

var (
	ErrBodyParsed   = errors.New("body could not be parsed")
	ErrCreatingNote = errors.New("fila creation Note")
)

type Handler struct {
	service notaService.Service
	log     logger.Logger
}

func New(log logger.Logger) *Handler {
	return &Handler{
		service: notaService.New(log),
		log:     log,
	}
}

type NoteCreateHandlerRequest struct {
	Title string `json:"titulo"`
	Body  string `json:"cuerpo"`
	Date  string `json:"fecha"`
}

func (h *Handler) NoteCreateHandler(w http.ResponseWriter, r *http.Request) {
	var tx NoteCreateHandlerRequest
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		h.log.Error(err)
		_ = response.HTTPError(w, r, http.StatusBadRequest, ErrBodyParsed.Error())
		return
	}
	key, err := h.service.NotaCreate(tx.Title, tx.Body, tx.Date)
	if err != nil {
		h.log.Error(err)
		_ = response.HTTPError(w, r, http.StatusUnprocessableEntity, ErrCreatingNote.Error())
	}

	defer r.Body.Close()
	_ = response.JSON(w, r, http.StatusOK, response.Map{"id": key})
}

func (h *Handler) NoteGetAllHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) NoteUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

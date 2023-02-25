package service

import (
	"netforemost/pkg/logger"
	"netforemost/pkg/nota/repository"
)

type Service interface {
	NotaCreate(title, body, date string) (int64, error)
	NotaGetAll(notaGetAllRequest NotaGetAllRequest) (notaGetAllResponseList []NotaGetAllResponse, err error)
	NotaUpdateById(id int64, title, body, date string) error
}

type service struct {
	repo repository.Repository
	log  logger.Logger
}

func New(log logger.Logger) Service {
	return &service{
		repo: repository.New(log),
		log:  log,
	}
}

func (s *service) NotaCreate(title, body, date string) (int64, error) {
	return s.repo.NotaCreate(repository.Nota{
		Title: title,
		Body:  body,
		Date:  date,
	})
}

type NotaGetAllRequest struct {
}

type NotaGetAllResponse struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
}

func (s *service) NotaGetAll(notaGetAllRequest NotaGetAllRequest) (notaGetAllResponseList []NotaGetAllResponse, err error) {
	notaList, err := s.repo.NotaGetAll("")
	if err != nil {
		return
	}
	for _, data := range notaList {
		notaGetAllResponseList = append(notaGetAllResponseList, NotaGetAllResponse{
			Id:    data.Id,
			Title: data.Title,
			Body:  data.Body,
			Date:  data.Date,
		})
	}

	return
}

func (s *service) NotaUpdateById(id int64, title, body, date string) error {
	return s.repo.NotaUpdate(repository.Nota{
		Id:    id,
		Title: title,
		Body:  body,
		Date:  date,
	})
}

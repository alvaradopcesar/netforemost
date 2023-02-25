package repository

import (
	"encoding/json"
	"strconv"

	"netforemost/pkg/cache"
	"netforemost/pkg/key_autoincremental"
	"netforemost/pkg/logger"
)

type Repository interface {
	NotaCreate(nota Nota) (id int64, err error)
	NotaGetAll(order string) (notaList []Nota, err error)
	NotaUpdate(nota Nota) (err error)
}

type repository struct {
	cache       cache.Cache
	log         logger.Logger
	incremental key_autoincremental.Incremental
}

func New(log logger.Logger) Repository {
	return &repository{
		cache:       cache.NewCache(),
		log:         log,
		incremental: key_autoincremental.New(),
	}
}

type Nota struct {
	Id    int64
	Title string
	Body  string
	Date  string
}

func (r *repository) NotaCreate(nota Nota) (id int64, err error) {
	id = r.incremental.Next()
	nota.Id = id
	b, err := json.Marshal(nota)
	if err != nil {
		return
	}
	err = r.cache.Set(strconv.FormatInt(id, 10), string(b), 0)
	if err != nil {
		return
	}
	return
}

func (r *repository) NotaGetAll(order string) (notaList []Nota, err error) {

	keys, err := r.cache.GetAllKeys()
	for _, notaDetail := range keys {
		nota := Nota{}
		err = json.Unmarshal([]byte(notaDetail), &nota)
		if err != nil {
			return notaList, err
		}
		notaList = append(notaList, nota)
	}
	return notaList, err
}

func (r *repository) NotaUpdate(nota Nota) (err error) {
	b, err := json.Marshal(nota)
	if err != nil {
		return
	}
	err = r.cache.Set(strconv.FormatInt(nota.Id, 10), string(b), 0)
	if err != nil {
		return
	}
	return
}

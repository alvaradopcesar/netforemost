package repository

import (
	"encoding/json"
	"reflect"
	"sort"
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
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
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
	sort := "id"
	switch order {
	case "id":
		sort = "id"
	case "titulo":
		sort = "title"
	case "cuerpo":
		sort = "body"
	case "fecha":
		sort = "date"
	}
	sortBy(sort, notaList)
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

func sortBy(jsonField string, arr []Nota) []Nota {
	if len(arr) < 1 {
		return []Nota{}
	}

	// first we find the field based on the json tag
	valueType := reflect.TypeOf(arr[0])

	var field reflect.StructField

	for i := 0; i < valueType.NumField(); i++ {
		field = valueType.Field(i)

		if field.Tag.Get("json") == jsonField {
			break
		}
	}

	// then we sort based on the type of the field
	sort.Slice(arr, func(i, j int) bool {
		v1 := reflect.ValueOf(arr[i]).FieldByName(field.Name)
		v2 := reflect.ValueOf(arr[j]).FieldByName(field.Name)

		switch field.Type.Name() {
		case "int":
			return int(v1.Int()) < int(v2.Int())
		case "string":
			return v1.String() < v2.String()
		case "bool":
			return !v1.Bool() // return small numbers first
		default:
			return false // return unmodified
		}
	})

	////fmt.Printf("\nsort by %s:\n", jsonField)
	//prettyPrint(arr)
	return arr
}

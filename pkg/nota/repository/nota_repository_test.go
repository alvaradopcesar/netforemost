package repository

import (
	"netforemost/pkg/cache"
	"netforemost/pkg/key_autoincremental"
	"netforemost/pkg/logger"
	"reflect"
	"testing"
)

func Test_repository_NotaCreate(t *testing.T) {
	ch := cache.NewCache()
	type fields struct {
		cache       cache.Cache
		log         logger.Logger
		incremental key_autoincremental.Incremental
	}
	type args struct {
		nota Nota
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "NoteCreate1",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			args: args{
				nota: Nota{
					Title: "Title01",
					Body:  "Body01",
					Date:  "2023/01/01",
				},
			},
			wantId:  1,
			wantErr: false,
		},
		{
			name: "NoteCreate2",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			args: args{
				nota: Nota{
					Title: "NoteCreate1",
					Body:  "Body02",
					Date:  "2023/01/02",
				},
			},
			wantId:  2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				cache:       tt.fields.cache,
				log:         tt.fields.log,
				incremental: tt.fields.incremental,
			}
			gotId, err := r.NotaCreate(tt.args.nota)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotaCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("NotaCreate() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}

}

func Test_repository_NotaGetAll(t *testing.T) {
	ch := cache.NewCache()
	type fields struct {
		cache       cache.Cache
		log         logger.Logger
		incremental key_autoincremental.Incremental
	}
	type args struct {
		order string
	}
	type argsCreate struct {
		nota Nota
	}
	tests := []struct {
		name         string
		fields       fields
		argsCreate   argsCreate
		args         args
		wantNotaList []Nota
		wantErr      bool
	}{
		{
			name: "NoteGetAll_id",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			argsCreate: argsCreate{
				nota: Nota{
					Title: "C Title",
					Body:  "C Body",
					Date:  "C Data",
				},
			},
			args: args{
				order: "id",
			},
			wantNotaList: []Nota{
				{
					Id:    1,
					Title: "C Title",
					Body:  "C Body",
					Date:  "C Data",
				},
			},
			wantErr: false,
		},
		{
			name: "NoteGetAll_body",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			argsCreate: argsCreate{
				nota: Nota{
					Title: "A Title",
					Body:  "A Body",
					Date:  "A Data",
				},
			},
			args: args{
				order: "body",
			},
			wantNotaList: []Nota{
				{
					Id:    2,
					Title: "A Title",
					Body:  "A Body",
					Date:  "A Data",
				},
				{
					Id:    1,
					Title: "C Title",
					Body:  "C Body",
					Date:  "C Data",
				},
			},
			wantErr: false,
		},
		{
			name: "NoteGetAll_title",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			argsCreate: argsCreate{
				nota: Nota{
					Title: "B Title",
					Body:  "B Body",
					Date:  "B Data",
				},
			},
			args: args{
				order: "title",
			},
			wantNotaList: []Nota{
				{
					Id:    2,
					Title: "A Title",
					Body:  "A Body",
					Date:  "A Data",
				},
				{
					Id:    3,
					Title: "B Title",
					Body:  "B Body",
					Date:  "B Data",
				},
				{
					Id:    1,
					Title: "C Title",
					Body:  "C Body",
					Date:  "C Data",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				cache:       tt.fields.cache,
				log:         tt.fields.log,
				incremental: tt.fields.incremental,
			}
			r.NotaCreate(tt.argsCreate.nota)
			gotNotaList, err := r.NotaGetAll(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotaGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNotaList, tt.wantNotaList) {
				t.Errorf("NotaGetAll() gotNotaList = %v, want %v", gotNotaList, tt.wantNotaList)
			}
		})
	}
}

func Test_repository_NotaUpdate(t *testing.T) {
	ch := cache.NewCache()
	type fields struct {
		cache       cache.Cache
		log         logger.Logger
		incremental key_autoincremental.Incremental
	}
	type argsCreate struct {
		nota Nota
	}
	type args struct {
		nota Nota
	}
	tests := []struct {
		name       string
		fields     fields
		argsCreate argsCreate
		args       args
		wantErr    bool
	}{
		{
			name: "NoteUpdate1",
			fields: fields{
				cache:       ch,
				log:         logger.New("test01", false),
				incremental: key_autoincremental.New(),
			},
			argsCreate: argsCreate{
				nota: Nota{
					Title: "Title 02",
					Body:  "Body 02",
					Date:  "Data 02",
				},
			},
			args: args{
				nota: Nota{
					Id:    1,
					Title: "Title xx",
					Body:  "Body xx",
					Date:  "Data xx",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				cache:       tt.fields.cache,
				log:         tt.fields.log,
				incremental: tt.fields.incremental,
			}
			r.NotaCreate(tt.argsCreate.nota)
			if err := r.NotaUpdate(tt.args.nota); (err != nil) != tt.wantErr {
				t.Errorf("NotaUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

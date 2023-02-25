package service

import (
	"netforemost/pkg/logger"
	"reflect"
	"testing"
)

//func TestNew(t *testing.T) {
//	tests := []struct {
//		name string
//		want Service
//	}{
//		// TODO: Add test cases.
//		{
//			name : "New1",
//			want : Service(),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := New(tt.args.log); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("New() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_service_NotaCreate(t *testing.T) {
	type args struct {
		title string
		body  string
		date  string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "NotaCreate1",
			args: args{
				title: "title1",
				body:  "body1",
				date:  "date1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "NotaCreate2",
			args: args{
				title: "title1",
				body:  "body1",
				date:  "date1",
			},
			want:    2,
			wantErr: false,
		},
	}
	log := logger.New("xx", false)
	s := New(log)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.NotaCreate(tt.args.title, tt.args.body, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotaCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NotaCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_NotaGetAll(t *testing.T) {
	type args struct {
		notaGetAllRequest NotaGetAllRequest
	}
	tests := []struct {
		name                       string
		args                       args
		wantNotaGetAllResponseList []NotaGetAllResponse
		wantErr                    bool
	}{
		{
			name: "NotaGetAllAll1",
			args: args{},
		},
	}
	log := logger.New("xx", false)
	s := New(log)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNotaGetAllResponseList, err := s.NotaGetAll(tt.args.notaGetAllRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotaGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNotaGetAllResponseList, tt.wantNotaGetAllResponseList) {
				t.Errorf("NotaGetAll() gotNotaGetAllResponseList = %v, want %v", gotNotaGetAllResponseList, tt.wantNotaGetAllResponseList)
			}
		})
	}
}

func Test_service_NotaUpdateById(t *testing.T) {
	type args struct {
		id    int64
		title string
		body  string
		date  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: " NotaUpdate1",
			args: args{
				id:    1,
				title: "title1",
				body:  "body1",
				date:  "date1",
			},
		},
	}
	log := logger.New("xx", false)
	s := New(log)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.NotaUpdateById(tt.args.id, tt.args.title, tt.args.body, tt.args.date); (err != nil) != tt.wantErr {
				t.Errorf("NotaUpdateById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

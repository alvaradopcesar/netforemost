package service

import (
	"netforemost/pkg/logger"
	"testing"
)

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

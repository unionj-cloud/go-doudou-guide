package service

import (
	"context"
	"os"
	"reflect"
	"testing"
	"usersvc/config"
)

func TestUsersvcImpl_DownloadAvatar(t *testing.T) {
	type fields struct {
		conf config.Config
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile *os.File
		wantErr  bool
	}{
		{
			name: "",
			fields: fields{
				conf: config.Config{},
			},
			args: args{
				ctx:    context.Background(),
				userId: "111",
			},
			wantFile: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &UsersvcImpl{
				conf: tt.fields.conf,
			}
			gotFile, err := receiver.DownloadAvatar(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("DownloadAvatar() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

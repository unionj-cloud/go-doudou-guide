package client

import (
	"context"
	"mime/multipart"
	"os"
	"reflect"
	"strings"
	"testing"
)

func init() {
	os.Setenv("USERSVC", "http://localhost:6060")
}

const (
	fileaContents = "This is a test file."
	filebContents = "Another test file."
	textaValue    = "foo"
	textbValue    = "bar"
	boundary      = `MyBoundary`
)

const message = `
--MyBoundary
Content-Disposition: form-data; name="file"; filename="filea.txt"
Content-Type: text/plain

` + fileaContents + `
--MyBoundary
Content-Disposition: form-data; name="file"; filename="fileb.txt"
Content-Type: text/plain

` + filebContents + `
--MyBoundary
Content-Disposition: form-data; name="texta"

` + textaValue + `
--MyBoundary
Content-Disposition: form-data; name="textb"

` + textbValue + `
--MyBoundary--
`

func TestUsersvcClient_UploadAvatar(t *testing.T) {
	type args struct {
		ctx     context.Context
		headers []*multipart.FileHeader
		s       string
	}
	b := strings.NewReader(strings.ReplaceAll(message, "\n", "\r\n"))
	r := multipart.NewReader(b, boundary)
	f, err := r.ReadForm(25)
	if err != nil {
		t.Fatal("ReadForm:", err)
	}
	defer f.RemoveAll()

	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				ctx:     context.Background(),
				headers: f.File["file"],
				s:       f.Value["texta"][0],
			},
			want:    0,
			want1:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUsersvc()
			got, got1, err := u.UploadAvatar(tt.args.ctx, tt.args.headers, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UploadAvatar() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UploadAvatar() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUsersvcClient_DownloadAvatar(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{
			name: "2",
			args: args{
				ctx:    context.Background(),
				userId: "2",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUsersvc()
			got, err := u.DownloadAvatar(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

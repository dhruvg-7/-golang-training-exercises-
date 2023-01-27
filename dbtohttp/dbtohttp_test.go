package dbtohttp

import (
	"database/sql"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_readDB(t *testing.T) {
	type args struct {
		myDb *sql.DB
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tests := []struct {
		name    string
		args    args
		want    []Person
		wantErr bool
	}{
		{
			name:    "Base case",
			args:    args{myDb: db},
			want:    []Person{},
			wantErr: false,
		},
	}
	mock.ExpectQuery("SELECT .* FROM person").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"Name", "0", "Phone"}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDB(tt.args.myDb)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_roothandlerPing(t *testing.T) {

// 	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
// 	w := httptest.NewRecorder()
// 	rootHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// 	if string("Pong") != string(data) {
// 		t.Errorf("expected Pong got %v", string(data))
// 	}

// }
// func Test_roothandlerPerson(t *testing.T) {

// 	req := httptest.NewRequest(http.MethodGet, "/person", nil)
// 	w := httptest.NewRecorder()
// 	rootHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 	}
// 	if !reflect.DeepEqual("", string(data)) { //sql data
// 		t.Errorf("expected sql data got %v", string(data))
// 	}

// }
func Test_roothandler(t *testing.T) {
	type args struct {
		target string
		reader io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{

			name: "RootHandlerPing",
			args: args{
				target: "/ping",
				reader: nil,
			},
			want:    "Pong",
			wantErr: false,
		},
		{

			name: "RootHandlerPerson",
			args: args{
				target: "/person",
				reader: nil,
			},
			want:    "", //sql data
			wantErr: false,
		},
		{

			name: "Root /",
			args: args{
				target: "/",
				reader: nil,
			},
			want:    "404",
			wantErr: false,
		},
		{

			name: "Root random",
			args: args{
				target: "/rand",
				reader: nil,
			},
			want:    "rand",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.args.target, tt.args.reader)
			w := httptest.NewRecorder()
			rootHandler(w, req)
			got := w.Result()

			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("rootHandler() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

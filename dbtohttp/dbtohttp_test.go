package dbtohttp

import (
	"database/sql"
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
		// TODO: Add test cases.
	}
	mock.ExpectQuery("select * from student").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Age", "Address"}))

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

func Test_personhandler() {
	testServer := httptest.NewRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

}

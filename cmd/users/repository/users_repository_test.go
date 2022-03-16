package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"regexp"
	"testing"
)

func TestNewUsersRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want UsersRepository
	}{
		{
			name: "Test with nil db should pass",
			args: args{db: nil},
			want: NewUsersRepository(nil),
		},
		{
			name: "Test with no nil db should pass",
			args: args{db: &gorm.DB{}},
			want: NewUsersRepository(&gorm.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsersRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsersRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersRepository_GetUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Error(err.Error())
	}
	gormDb.Debug()

	row := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL")).
		WillReturnRows(row)

	userRepository := NewUsersRepository(gormDb)

	result, err := userRepository.GetUser("email@abc.com")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUsersRepository_GetUser_Error(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Error(err.Error())
	}
	gormDb.Debug()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("mocked error from db"))

	userRepository := NewUsersRepository(gormDb)

	result, err := userRepository.GetUser("email@abc.com")

	assert.Error(t, err)
	assert.Nil(t, result)
}

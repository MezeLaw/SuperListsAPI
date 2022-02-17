package repository

import (
	"SuperListsAPI/cmd/auth/models"
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

func TestNewAuthRepository(t *testing.T) {

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want AuthRepository
	}{
		{
			name: "Test with nil db should create a repo with nil db",
			args: args{db: nil},
			want: NewAuthRepository(nil),
		},
		{
			name: "Test with no nil db should create a repo with no nil db",
			args: args{db: &gorm.DB{}},
			want: NewAuthRepository(&gorm.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthRepository_SignUp(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validUser := GetValidUser()
	validUser.Role = ""

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := authRepository.SignUp(validUser)

	assert.NotEmpty(t, result)

}

func TestAuthRepository_SignUp_Error(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validUser := &models.User{}
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users")).
		WillReturnError(errors.New("error when insert into users"))
	mock.ExpectCommit()

	result, err := authRepository.SignUp(validUser)

	assert.Empty(t, result)
	assert.Error(t, err)

}

func TestAuthRepository_Login(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validLogin := GetValidLoginPayload()

	validUser := GetValidUser()

	validUser.HashPassword(validLogin.Password)

	hashedPass := validUser.Password

	row := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "Meze", "meze@meze.com", hashedPass)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(row)

	result, err := authRepository.Login(*validLogin)

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}

func TestAuthRepository_Login_Invalid_Credentials(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validLogin := GetValidLoginPayload()

	validUser := GetValidUser()

	validUser.HashPassword(validLogin.Password)

	row := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "Meze", "meze@meze.com", "invalidPassword")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(row)

	result, err := authRepository.Login(*validLogin)

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestAuthRepository_Login_Email_Not_Found(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validLogin := GetValidLoginPayload()

	validUser := GetValidUser()

	validUser.HashPassword(validLogin.Password)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnError(errors.New("record not found"))

	result, err := authRepository.Login(*validLogin)

	assert.Empty(t, result)
	assert.Error(t, err)
}

func TestAuthRepository_Login_Email_Not_Found_Generic_Error(t *testing.T) {
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

	authRepository := NewAuthRepository(gormDb)

	validLogin := GetValidLoginPayload()

	validUser := GetValidUser()

	validUser.HashPassword(validLogin.Password)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`"))

	result, err := authRepository.Login(*validLogin)

	assert.Empty(t, result)
	assert.Error(t, err)
}

func GetValidLoginPayload() *models.LoginPayload {
	return &models.LoginPayload{
		Email:    "meze@meze.com",
		Password: "password",
	}
}

func GetValidUser() *models.User {
	return &models.User{
		Name:     "Meze",
		Email:    "meze@meze.com",
		Password: "password",
		Role:     "ADMIN",
	}
}

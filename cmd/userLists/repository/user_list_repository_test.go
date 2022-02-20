package repository

import (
	"SuperListsAPI/cmd/userLists/models"
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

func TestNewUserListRepository(t *testing.T) {
	type args struct {
		gormDB *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want UserListRepository
	}{
		{
			name: "Test with nil gormDB should pass",
			args: args{gormDB: nil},
			want: NewUserListRepository(nil),
		},
		{
			name: "Test with no nil gormDb should pass",
			args: args{gormDB: &gorm.DB{}},
			want: NewUserListRepository(&gorm.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserListRepository(tt.args.gormDB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserListRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserListRepository_Create(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user_lists`" +
		" (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`) VALUES (?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.Create(GetValidUserList())

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestUserListRepository_Create_Error(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user_lists`" +
		" (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`) VALUES (?,?,?,?,?)")).
		WillReturnError(errors.New("error from db"))
	mock.ExpectCommit()

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.Create(GetValidUserList())

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestUserListRepository_Get(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "list_id"}).AddRow(1, 1))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.Get("1")

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
}

func TestUserListRepository_Get_Error(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error from db on get method"))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.Get("1")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestUserListRepository_Delete(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `user_lists` SET `deleted_at`=? " +
		"WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userListRepository := NewUserListRepository(gormDb)

	result, err := userListRepository.Delete("1")

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, *result, 1)

}

func TestUserListRepository_Delete_Error(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `user_lists` SET `deleted_at`=? " +
		"WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error trying to delete user list"))
	mock.ExpectCommit()

	userListRepository := NewUserListRepository(gormDb)

	result, err := userListRepository.Delete("1")

	assert.Nil(t, result)
	assert.Error(t, err)

}

func TestUserListRepository_GetUserListsByUserID(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE user_id = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "list_id"}).AddRow(1, 1).AddRow(2, 2))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.GetUserListsByUserID("1")

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestUserListRepository_GetUserListsByUserID_Error(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error from db on getUserLists by userID method"))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.GetUserListsByUserID("1")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestUserListRepository_GetUserListsByListID(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE list_id = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "list_id"}).AddRow(1, 1).AddRow(2, 2))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.GetUserListsByListID("1")

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestUserListRepository_GetUserListsByListID_Error(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE `user_lists`.`id` = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error from db on getUserLists by userID method"))

	userListRepo := NewUserListRepository(gormDb)

	result, err := userListRepo.GetUserListsByListID("1")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func GetValidUserList() models.UserList {
	return models.UserList{
		Model:  gorm.Model{},
		ListID: 1,
		UserID: 1,
	}
}

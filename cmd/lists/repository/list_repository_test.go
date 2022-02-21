package repository

import (
	"SuperListsAPI/cmd/lists/models"
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

func TestNewListRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want ListRepository
	}{
		{
			name: "Test with nil db should pass",
			args: args{db: nil},
			want: NewListRepository(nil),
		},
		{
			name: "Test with mocked repo should pass",
			args: args{db: &gorm.DB{}},
			want: NewListRepository(&gorm.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListRepository_Create(t *testing.T) {
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

	validList := GetValidList()

	listRepository := NewListRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `lists` (`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`invite_code`,`user_creator_id`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	//WillReturnError(errors.New("error when insert into lists"))
	mock.ExpectCommit()

	result, err := listRepository.Create(validList)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestListRepository_Create_Error(t *testing.T) {
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

	validList := GetValidList()

	listRepository := NewListRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `lists` (`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`invite_code`,`user_creator_id`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnError(errors.New("error when insert into lists"))
	mock.ExpectCommit()

	result, err := listRepository.Create(validList)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_GetLists(t *testing.T) {

	//gormDb, mock := GetMockDB()

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE user_id = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnRows(row)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `lists` WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL")).
		WillReturnRows(row)

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.GetLists("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListRepository_GetLists_List_Repo_Error(t *testing.T) {

	//gormDb, mock := GetMockDB()

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `user_lists` WHERE user_id = ? AND `user_lists`.`deleted_at` IS NULL")).
		WillReturnRows(row)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `lists` WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error from list db"))

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.GetLists("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_GetLists_Error(t *testing.T) {

	//gormDb, mock := GetMockDB()

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `lists` WHERE (user_creator_id = ?) AND `lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error when trying to get lists"))

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.GetLists("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_Get(t *testing.T) {

	//gormDb, mock := GetMockDB()

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `lists` WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL ORDER BY `lists`.`id` LIMIT 1")).
		WillReturnRows(row)

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.Get("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListRepository_Get_Error(t *testing.T) {
	//TODO mejorar e ir a esta forma
	//gormDb, mock := GetMockDB()

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `lists` WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL ORDER BY `lists`.`id` LIMIT 1")).
		WillReturnError(errors.New("error when trying to get lists"))

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.Get("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_Update(t *testing.T) {
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

	validList := GetValidList()

	listRepository := NewListRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `lists` (`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`invite_code`,`user_creator_id`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := listRepository.Update(validList)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.ID, uint(1))
}

func TestListRepository_Update_Error(t *testing.T) {
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

	validList := GetValidList()

	listRepository := NewListRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `lists` (`created_at`,`updated_at`,`deleted_at`,`name`,`description`,`invite_code`,`user_creator_id`) VALUES (?,?,?,?,?,?,?)")).
		WillReturnError(errors.New("error when updating into lists"))
	mock.ExpectCommit()

	result, err := listRepository.Update(validList)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_Delete(t *testing.T) {
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
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `lists` SET `deleted_at`=? WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.Delete("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListRepository_Delete_Error(t *testing.T) {
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
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `lists` SET `deleted_at`=? WHERE `lists`.`id` = ? AND `lists`.`deleted_at` IS NULL")).
		WillReturnError(errors.New("error when updating list"))
	mock.ExpectCommit()

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.Delete("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListRepository_GetListByInvitationCode(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.GetListByInvitationCode("mockCode")

	assert.NotNil(t, result)
	assert.NoError(t, err)

}

func TestListRepository_GetListByInvitationCode_Error(t *testing.T) {
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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnError(errors.New("error from list repository"))

	listRepository := NewListRepository(gormDb)

	result, err := listRepository.GetListByInvitationCode("mockCode")

	assert.Nil(t, result)
	assert.Error(t, err)

}

func GetValidList() models.List {
	return models.List{
		Model:         gorm.Model{},
		Name:          "mocked name",
		Description:   "mocked description",
		InviteCode:    "mockedCode",
		UserCreatorID: 1,
	}
}

func GetMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	var t *testing.T
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

	return gormDb, mock
}

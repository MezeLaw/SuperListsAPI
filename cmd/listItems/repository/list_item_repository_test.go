package repository

import (
	"SuperListsAPI/cmd/listItems/models"
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

func TestNewListItemRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want ListItemRepository
	}{
		{
			name: "Test with nil gormDB should pass",
			args: args{nil},
			want: NewListItemRepository(nil),
		},
		{
			name: "Test with no nil gormDB should pass",
			args: args{db: &gorm.DB{}},
			want: NewListItemRepository(&gorm.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListItemRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListItemRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListItemRepository_Create(t *testing.T) {

	validListItem := GetValidListItem()

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `list_items` (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`,`title`,`description`,`is_done`) " +
		"VALUES (?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := listItemRepository.Create(validListItem)

	assert.NotNil(t, result)
	assert.NoError(t, err)

}

func TestListItemRepository_Create_Error(t *testing.T) {

	validListItem := GetValidListItem()

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `list_items` (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`,`title`,`description`,`is_done`) " +
		"VALUES (?,?,?,?,?,?,?,?)")).
		WillReturnError(errors.New("Error from DB"))
	mock.ExpectCommit()

	result, err := listItemRepository.Create(validListItem)

	assert.Nil(t, result)
	assert.Error(t, err)

}

func TestListItemRepository_Update(t *testing.T) {

	validListItem := GetValidListItem()

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `list_items` (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`,`title`,`description`,`is_done`) VALUES (?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := listItemRepository.Update(validListItem)

	assert.NotNil(t, result)
	assert.NoError(t, err)

}

func TestListItemRepository_Update_Error(t *testing.T) {

	validListItem := GetValidListItem()

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `list_items` (`created_at`,`updated_at`,`deleted_at`,`list_id`,`user_id`,`title`,`description`,`is_done`) VALUES (?,?,?,?,?,?,?,?)")).
		WillReturnError(errors.New("Error from DB"))
	mock.ExpectCommit()

	result, err := listItemRepository.Update(validListItem)

	assert.Nil(t, result)
	assert.Error(t, err)

}

func TestListItemRepository_Get(t *testing.T) {

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `list_items` WHERE `list_items`.`id` = ? AND " +
		"`list_items`.`deleted_at` IS NULL ORDER BY `list_items`.`id` LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	result, err := listItemRepository.Get("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestListItemRepository_Get_Error(t *testing.T) {

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

	listItemRepository := NewListItemRepository(gormDb)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `list_items` WHERE `list_items`.`id` = ? AND " +
		"`list_items`.`deleted_at` IS NULL ORDER BY `list_items`.`id` LIMIT 1")).
		WillReturnError(errors.New("error from db"))

	result, err := listItemRepository.Get("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestListItemRepository_Delete(t *testing.T) {

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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := listItemRepo.Delete("1")

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestListItemRepository_Delete_Error(t *testing.T) {

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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WillReturnError(errors.New("error from db"))
	mock.ExpectCommit()

	result, err := listItemRepo.Delete("1")

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestListItemRepository_GetItemsListByListID(t *testing.T) {
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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

	result, err := listItemRepo.GetItemsListByListID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestListItemRepository_GetItemsListByListID_Error(t *testing.T) {
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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnError(errors.New("error from db"))

	result, err := listItemRepo.GetItemsListByListID("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestListItemRepository_DeleteListItemsByListID(t *testing.T) {

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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec((regexp.QuoteMeta("UPDATE"))).WillReturnResult(sqlmock.NewResult(1, 2))

	mock.ExpectCommit()

	result, err := listItemRepo.DeleteListItemsByListID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListItemRepository_DeleteListItemsByListID_Error(t *testing.T) {

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

	listItemRepo := NewListItemRepository(gormDb)

	mock.ExpectBegin()
	mock.ExpectExec((regexp.QuoteMeta("UPDATE"))).
		WillReturnError(errors.New("error from db"))

	mock.ExpectCommit()

	result, err := listItemRepo.DeleteListItemsByListID("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func GetValidListItem() models.ListItem {
	return models.ListItem{
		ListID:      1,
		UserID:      1,
		Title:       "Hacer la tarea",
		Description: "Completar las funciones cuadraticas, graficarlas y derivarlas",
		IsDone:      false,
	}
}

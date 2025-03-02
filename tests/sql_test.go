package null_test

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/ogiusek/null"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbFile string = "./test_db.db"

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Model{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Save(db *gorm.DB, instance Model) error {
	tx := db.Save(&instance)
	return tx.Error
}

func Delete(db *gorm.DB, instance Model) error {
	tx := db.Where("id = ?", instance.Id).Delete(&Model{})
	return tx.Error
}

func Get(db *gorm.DB, id string) *Model {
	var model Model
	tx := db.Where("id = ?", id).First(&model)
	if tx.Error != nil {
		log.Print("error", tx.Error)
		return nil
	}

	return &model
}

func Clear(db *gorm.DB) error {
	tx := db.Where("1 = 1").Delete(&Model{})
	return tx.Error
}

type Model struct {
	Id    string                `gorm:"column:id;primaryKey;type:VARCHAR(36)"`
	Value null.Nullable[string] `gorm:"column:value;null"`
}

func NewModel(value null.Nullable[string]) Model {
	return Model{
		Id:    uuid.NewString(),
		Value: value,
	}
}

//

func TestOptionalInDB(t *testing.T) {
	db := Connection()
	if err := Clear(db); err != nil {
		t.Errorf("error occured when trying to clear db %s", err.Error())
		return
	}
	var instances []Model = []Model{
		NewModel(null.Null[string]()),
		NewModel(null.New("")),
		NewModel(null.New("a")),
		NewModel(null.New("b")),
		NewModel(null.New("c")),
	}

	for _, instance := range instances {
		if err := Save(db, instance); err != nil {
			t.Errorf("error occured saving model %s\n%v", err.Error(), instance)
			return
		}

		retrieved := Get(db, instance.Id)
		if retrieved == nil {
			t.Errorf("retrieved model should exist and be %v", instance)
		} else if *retrieved != instance {
			t.Errorf("\nexpected  %v\nretrieved %v", instance, *retrieved)
		}

		if err := Delete(db, instance); err != nil {
			t.Errorf("error occured deleting model %s\n%v", err.Error(), instance)
			return
		}
	}
}

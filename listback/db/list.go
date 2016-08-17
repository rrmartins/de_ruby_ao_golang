package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Listing struct {
	gorm.Model
	Title       string
	Description string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	checkErr(err)
	db.AutoMigrate(&Listing{})
}

func Insert(title, description string) {
	listing := Listing{Title: title, Description: description}
	db.NewRecord(listing)
	db.Create(&listing)
}

func Update(id int, title, description string) {
	var listing Listing
	db.First(&listing, "id = ?", id)
	listing.Title = title
	listing.Description = description
	db.Save(&listing)
}

func ListAll() []Listing {
	var listings []Listing
	db.Find(&listings)

	return listings
}

func Delete(id int) {
	db.Where("id = ?", &id).Delete(&Listing{})
}

func Show(id int) Listing {
	var listing Listing
	db.First(&listing, "id = ?", id)
	return listing
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

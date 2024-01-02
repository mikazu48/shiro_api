package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/shiro_db?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConn(t *testing.T) {
	assert.NotNil(t, db)
}

func TestInsert(t *testing.T) {
	err := db.Exec("insert into user(username, password, email, full_name) values (?, ?, ?, ?) ", "kazu", "kazu", "kazu@gg.gg", "Dimas Uhuy").Error
	assert.Nil(t, err)
}

func TestPassword(t *testing.T) {
	password := "cipeng123!!"

	hashedPassword := "$2a$10$X8dDbqefWzdz54Wv8nWg7eqNleEod8UUk3OLJuOuo6DcYbQABX5NO"
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	require.NoError(t, err)

	// wrongPassword := "aw"
	// err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(wrongPassword))
	// require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}

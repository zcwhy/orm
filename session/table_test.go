package session

import (
	"database/sql"
	"orm/dialect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// CREATE TABLE `User` (`Name` text PRIMARY KEY, `Age` integer);
type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_Model(t *testing.T) {
	s := NewTestSession()
	s.Model(&Session{})
	if s.refSchema.Name != "Session" {
		t.Fatal("Failed to set model")
	}

	s.Model(&User{})
	if s.refSchema.Name != "User" {
		t.Fatal("Failed to change model")
	}
}

func TestSession_CreateTable(t *testing.T) {
	s := NewTestSession().Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
}

func NewTestSession() *Session {
	testDB, _ := sql.Open("sqlite3", "../gee.db")
	testDial, _ := dialect.GetDialector("sqlite3")

	return New(testDB, testDial)
}

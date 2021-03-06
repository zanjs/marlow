package models

//go:generate marlowc -input author.go

// Example file for testing.

// Author represents an author of a book.
type Author struct {
	_    bool   `marlow:"tableName=authors"`
	ID   int    `marlow:"column=id"`
	Name string `marlow:"column=name"`
}

package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

const (
	INTEGER_OBJ = "INTEGER"
)

// Type returns the type of the object
func (i Integer) Type() ObjectType { return INTEGER_OBJ }

// Inspect returns a stringified version of the object for debugging
func (i Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

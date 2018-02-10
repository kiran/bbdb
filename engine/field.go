package engine

import (
	"bufio"
	"strconv"
)

// Field is the interface for values of fields in tuples.
type Field interface {
	GetType() Type
	// serialize writes the bytes representing this field to the specified
	// bufio.Writer
	serialize(writeout *bufio.Writer) error

	String() string

	// todo: compare
}

type IntField struct {
	value int
}

func (intf IntField) String() string {
	return strconv.Itoa(intf.value)
}

func (intf IntField) GetType() Type {
	return IntType{}
}

func (intf IntField) serialize(writeout *bufio.Writer) error {
	return nil // todo
}

type StringField struct {
	value string
}

func (stringf StringField) String() string {
	return stringf.value
}

func (stringf StringField) serialize(writeout *bufio.Writer) error {
	return nil // todo
}

func (stringf StringField) GetType() Type {
	return StringType{}
}

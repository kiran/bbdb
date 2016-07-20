package engine

import "bufio"

// Field is the interface for values of fields in tuples.
type Field interface {
	GetType() Type
	// serialize writes the bytes representing this field to the specified
	// bufio.Writer
	serialize(writeout *bufio.Writer) error

	// todo: compare
}

type IntField struct {
	value int
}

func (intf *IntField) serialize(writeout *bufio.Writer) error {
	return nil // todo
}

type StringField struct {
	value string
}

func (stringf *StringField) serialize(writeout *bufio.Writer) error {
	return nil // todo
}

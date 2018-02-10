package engine

import "bufio"

// Type represents a type in bbdb.
// Types are static objects defined by this struct, so you can't create a new
// type (e.g., it has no New)
type Type interface {
	// Length returns the number of bytes required to store a field of this type
	Length() int

	// Parse returns a Field object of the same type as this object that has
	// contents read from the specified data input stream.
	Parse(dis *bufio.Reader) (Field, error)

	String() string
}

// define type int and type string
type IntType struct{}

func (it IntType) Length() int {
	return 4
}

// TODO
func (it IntType) Parse(dis *bufio.Reader) (Field, error) {
	return nil, nil
}

func (it IntType) String() string {
	return "IntType"
}

type StringType struct{}

func (st StringType) Length() int {
	// TODO
	return 12
}

func (st StringType) Parse(dis *bufio.Reader) (Field, error) {
	return nil, nil
}

func (st StringType) String() string {
	return "StringType"
}

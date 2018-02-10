package engine

import (
	"bytes"
	"errors"
	"fmt"
)

// TupleDesc describes the schema of a Tuple.
type TupleDesc struct {
	typeAr  []Type
	fieldAr []string
}

// Combine merges two TupleDescs into one, with td1.numFields + td2.numfields
// fields, with the first fields from td1, and the remaining from td2.
func Combine(td1 *TupleDesc, td2 *TupleDesc) (*TupleDesc, error) {
	if td1 == nil || td2 == nil {
		return nil, errors.New("can't have td1 or td2 be nil")
	}
	// make a new type ar and a new field ar, then initialize a new tupledesc
	newTypeAr := append(td1.typeAr, td2.typeAr...)
	newFieldAr := append(td1.fieldAr, td2.fieldAr...)

	return &TupleDesc{typeAr: newTypeAr, fieldAr: newFieldAr}, nil
}

// Create a new TupleDesc with fields of the specified type, with associated
// name fields.
func NewTupleDesc(typeAr []Type, fieldAr []string) (*TupleDesc, error) {
	if len(typeAr) != len(fieldAr) {
		return nil, errors.New(
			fmt.Sprintf(
				"typeAr and fieldAr must be of the same length; got %d and %d",
				len(typeAr), len(fieldAr)))
	}
	td := TupleDesc{typeAr: typeAr, fieldAr: fieldAr}
	return &td, nil
}

// Create a new tuple desc with typeAr.length fields with fields of the
// specified types, with anonymous (unnamed) fields.
func NewAnonTupleDesc(typeAr []Type) *TupleDesc {
	return &TupleDesc{typeAr: typeAr}
}

// NumFields returns the number of fields in this TupleDesc
func (td *TupleDesc) NumFields() int {
	return len(td.typeAr)
}

// GetFieldName gets the (possibly null) field name of the ith field of this
// TupleDesc
func (td *TupleDesc) GetFieldName(i int) (string, error) {
	if i >= len(td.fieldAr) {
		return "", errors.New("No such element")
	}
	return td.fieldAr[i], nil
}

// NameToID finds the index of the field with a given name
func (td *TupleDesc) NameToID(name string) (int, error) {
	for i, fieldname := range td.fieldAr {
		if fieldname == name {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
}

// GetType returns the type of the ith field of this TupleDesc
func (td *TupleDesc) GetType(i int) (Type, error) {
	if i >= len(td.typeAr) {
		return nil, errors.New("No such element")
	}
	return td.typeAr[i], nil
}

// GetSize returns the size (in bytes) of tuples corresponding to this TupleDesc
// Note that tuples from a given TupleDesc are of a fixed size.
func (td *TupleDesc) ByteSize() int {
	// TODO
	return 0
}

// String returns a string describing this descriptor.
// "fieldType[0](fieldName[0]), ..., fieldType[M](fieldName[M])"
func (td *TupleDesc) String() string {
	var buffer bytes.Buffer

	for i, typ := range td.typeAr {
		if typ != nil {
			buffer.WriteString(typ.String())
		}
		// need to check fieldAr length
		if len(td.fieldAr) != 0 {
			buffer.WriteRune('(')
			buffer.WriteString(td.fieldAr[i])
			buffer.WriteRune(')')
		}
		buffer.WriteString(", ")
	}

	buffer.WriteString("\n")
	return buffer.String()
}

// TODO
// do I need to write an equals method?
// how does toString work in Go?

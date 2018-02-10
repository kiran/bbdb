package engine

import (
	"bytes"
	"errors"
)

// Tuple maintains information about the contents of a tuple.
// Tuples have a schema specified by a Tuple object and contain Field
// objects with the data for each Field.
type Tuple struct {
	desc   *TupleDesc
	rid    *RecordID
	fields []Field
}

// NewTuple creates a new tuple with the specified schema. The argument
// must be a valid TupleDesc with at least one field.
func NewTuple(td *TupleDesc) *Tuple {
	// verify that TupleDesc has at least one field.
	tuple := Tuple{desc: td}
	tuple.fields = make([]Field, td.NumFields())
	return &tuple
}

func (tup *Tuple) TupleDesc() *TupleDesc {
	return tup.desc
}

// the RecordID representing the location of this tuple on disk.
// May be nil.
func (tup *Tuple) RecordID() *RecordID {
	return tup.rid
}

func (tup *Tuple) SetRecordID(rid *RecordID) {
	tup.rid = rid
}

// SetField changes the value of the ith field of this tuple to f
func (tup *Tuple) SetField(i int, f Field) error {
	if i >= len(tup.fields) {
		return errors.New("out of bounds!")
	}
	tup.fields[i] = f
	return nil
}

// return the value of the ith field (may be null if not set)
func (tup *Tuple) Field(i int) (Field, error) {
	if i >= len(tup.fields) {
		return nil, errors.New("out of bounds!")
	}
	return tup.fields[i], nil
}

// represents string as: column1\tcolumn2\t...\tcolumnN\n
func (tup *Tuple) String() string {
	var buffer bytes.Buffer

	for _, field := range tup.fields {
		if field != nil {
			buffer.WriteString(field.String())
		}
		buffer.WriteString("\t")
	}

	buffer.WriteString("\n")
	return buffer.String()
}

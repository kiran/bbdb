package engine

// Tuple maintains information about the contents of a tuple.
// Tuples have a schema specified by a TupleDesc object and contain Field
// objects with the data for each Field.
type Tuple struct {
	desc *TupleDesc
}

// NewTuple creates a new tuple with the specified schema. The argument
// must be a valid TupleDesc with at least one field.
func NewTuple(td *TupleDesc) *Tuple {
	//todo
	return nil
}

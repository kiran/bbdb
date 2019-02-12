package engine

import "testing"

func TestModifyFields(t *testing.T) {
	typeAr := []Type{&IntType{}, &IntType{}}
	td := NewAnonTupleDesc(typeAr)

	tup := NewTuple(td)

	// test getting nil val
	got, _ := tup.Field(0)
	if got != nil {
		t.Errorf("tup.Field(0) = %d; want nil", got)
	}

	// test getting out-of-bounds val
	_, err := tup.Field(2)
	if err == nil {
		t.Errorf("expected this to error out")
	}

	tup.SetField(0, IntField{value: -1})
	tup.SetField(1, IntField{value: 2})

	got, _ = tup.Field(0)
	want := IntField{value: -1}
	if got != want {
		t.Errorf("tup.Field(0) = %d; want %d", got, want)
	}
	got, _ = tup.Field(1)
	want = IntField{value: 2}
	if got != want {
		t.Errorf("tup.Field(1) = %d; want %d", got, want)
	}

	// modify these fields
	tup.SetField(0, IntField{value: 1})
	tup.SetField(1, IntField{value: 37})

	got, _ = tup.Field(0)
	want = IntField{value: 1}
	if got != want {
		t.Errorf("tup.Field(0) = %d; want %d", got, want)
	}
	got, _ = tup.Field(1)
	want = IntField{value: 37}
	if got != want {
		t.Errorf("tup.Field(1) = %d; want %d", got, want)
	}
}

func TestGetTupleDesc(t *testing.T) {
	typeAr := []Type{&IntType{}, &IntType{}}
	td := NewAnonTupleDesc(typeAr)

	tup := NewTuple(td)

	got := tup.TupleDesc()
	if got != td {
		t.Errorf("Got an incorrect tuple desc")
	}
}

// add a test for modifying RecordID

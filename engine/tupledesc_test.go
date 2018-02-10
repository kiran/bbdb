package engine

import (
	"testing"
	)

func TestCombine(t *testing.T) {
	var td1, td2, td3 *TupleDesc
	// create 2 tds
	typeAr := []Type{&IntType{}, &IntType{}}
	td1, _ = NewTupleDesc(typeAr, []string{"td1", "td2"})
	typeAr = []Type{&StringType{}, &IntType{}}
	td2, _ = NewTupleDesc(typeAr, []string{"td3", "td4"})
	
	// test combine(td1, td2)
	td3, err := Combine(td1, td2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if td3.NumFields() != 4 {
		t.Errorf("Got %d fields, want %d", td3.NumFields(), 4)
	}
	expected_size := 3*IntType{}.Length() + StringType{}.Length()
	if td3.ByteSize() != expected_size {
		t.Errorf("Got %d as byte size, want %d", td3.ByteSize(), expected_size)
	}

	// test combine(td2, td1)

}

// func TestReverse(t *testing.T) {
// 	cases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{"Hello, 世界", "界世 ,olleH"},
// 		{"", ""},
// 	}
// 	for _, c := range cases {
// 		got := Reverse(c.in)
// 		if got != c.want {
// 			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
// 		}
// 	}
// }

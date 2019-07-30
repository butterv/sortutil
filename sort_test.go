package sortutil

import "testing"

func TestSort_Asc(t *testing.T) {
	in := []struct {
		ID    uint64
		Name  string
		Level uint
	}{
		{
			ID:    1,
			Name:  "enemy1",
			Level: 1,
		},
		{
			ID:    2,
			Name:  "enemy2",
			Level: 2,
		},
		{
			ID:    3,
			Name:  "enemy3",
			Level: 3,
		},
	}

	_, err := New(in).Order("ID", ASC)
	if err != nil {
		t.Fatal(err.Error())
	}

	is := [3]int{1, 2, 3}
	_, err = New(is).Order("ID", DESC)
	if err != nil {
		t.Fatal(err.Error())
	}
}

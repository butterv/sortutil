package sortutil

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Test struct {
	ID        uint64
	Capital   string
	Latitude  float32
	Longitude float32
}

var tests = []Test{
	{
		ID:        3,
		Capital:   "Tokyo",
		Latitude:  35.67581,
		Longitude: 139.74507,
	},
	{
		ID:        5,
		Capital:   "Seoul",
		Latitude:  37.531084,
		Longitude: 126.91583,
	},
	{
		ID:        2,
		Capital:   "Beijing",
		Latitude:  39.908222,
		Longitude: 116.391027,
	},
	{
		ID:        1,
		Capital:   "Washington D.C.",
		Latitude:  38.897159,
		Longitude: -77.036207,
	},
	{
		ID:        4,
		Capital:   "London",
		Latitude:  51.499183,
		Longitude: -0.12464066,
	},
	{
		ID:        6,
		Capital:   "Brasilia",
		Latitude:  -15.799668,
		Longitude: -47.864154,
	},
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func TestSort_Order_Asc(t *testing.T) {
	tempTests := make([]Test, len(tests))
	copy(tempTests, tests)

	t.Run("order by ID asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.ID)
		}
		New(tempTests).Order("ID", ASC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.ID)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].ID < tempTests[j].ID
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Capital asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Capital)
		}
		New(tempTests).Order("Capital", ASC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Capital)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Capital < tempTests[j].Capital
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Latitude asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Latitude)
		}
		New(tempTests).Order("Latitude", ASC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Latitude)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Latitude < tempTests[j].Latitude
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Longitude asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Longitude)
		}
		New(tempTests).Order("Longitude", ASC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Longitude)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Longitude < tempTests[j].Longitude
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})
}

func TestSort_Order_Desc(t *testing.T) {
	tempTests := make([]Test, len(tests))
	copy(tempTests, tests)

	t.Run("order by ID desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.ID)
		}
		New(tempTests).Order("ID", DESC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.ID)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].ID > tempTests[j].ID
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Capital desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Capital)
		}
		New(tempTests).Order("Capital", DESC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Capital)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Capital > tempTests[j].Capital
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Latitude desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Latitude)
		}
		New(tempTests).Order("Latitude", DESC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Latitude)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Latitude > tempTests[j].Latitude
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Longitude desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Longitude)
		}
		New(tempTests).Order("Longitude", DESC)
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%fs, ", test.Longitude)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Longitude > tempTests[j].Longitude
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkSort_Order_By_ID_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("ID", ASC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].ID < tempTests[j].ID
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Capital_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Capital", ASC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Capital < tempTests[j].Capital
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Latitude_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Latitude", ASC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Latitude < tempTests[j].Latitude
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Longitude_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Longitude", ASC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Longitude < tempTests[j].Longitude
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_ID_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("ID", DESC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].ID > tempTests[j].ID
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Capital_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Capital", DESC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Capital > tempTests[j].Capital
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Latitude_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Latitude", DESC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Latitude > tempTests[j].Latitude
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Longitude_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:        rand.Uint64(),
			Capital:   randStringRunes(8),
			Latitude:  rand.Float32(),
			Longitude: rand.Float32(),
		})
	}

	b.ResetTimer()
	New(tempTests).Order("Longitude", DESC)

	sortFunc := func(i, j int) bool {
		return tempTests[i].Longitude > tempTests[j].Longitude
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

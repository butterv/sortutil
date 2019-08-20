package sortutil

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Test struct {
	ID     uint64
	Name   string
	Age    uint
	Height float32
	Weight float32
}

var tests = []Test{
	{
		ID:     1,
		Name:   "Tom",
		Age:    20,
		Height: 170.5,
		Weight: 70.0,
	},
	{
		ID:     7,
		Name:   "Bob",
		Age:    30,
		Height: 180.5,
		Weight: 80.0,
	},
	{
		ID:     5,
		Name:   "Zach",
		Age:    30,
		Height: 175.0,
		Weight: 75.0,
	},
	{
		ID:     3,
		Name:   "Angie",
		Age:    20,
		Height: 160.5,
		Weight: 55.0,
	},
	{
		ID:     6,
		Name:   "Sam",
		Age:    25,
		Height: 178.5,
		Weight: 73.0,
	},
	{
		ID:     2,
		Name:   "Taro",
		Age:    18,
		Height: 179.5,
		Weight: 70.0,
	},
	{
		ID:     4,
		Name:   "Kim",
		Age:    25,
		Height: 150.2,
		Weight: 48.0,
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

		Order(tempTests).Asc("ID").Exec()
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

	t.Run("order by Name asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Name)
		}
		Order(tempTests).Asc("Name").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Name)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Name < tempTests[j].Name
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Height asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Height)
		}
		Order(tempTests).Asc("Height").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Height)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Height < tempTests[j].Height
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Age asc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.Age)
		}
		Order(tempTests).Asc("Age").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%d, ", test.Age)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Age < tempTests[j].Age
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
		Order(tempTests).Desc("ID").Exec()
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

	t.Run("order by Name desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Name)
		}
		Order(tempTests).Desc("Name").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%s, ", test.Name)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Name > tempTests[j].Name
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Weight desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Weight)
		}
		Order(tempTests).Desc("Weight").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Weight)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Weight > tempTests[j].Weight
		}
		if !sort.SliceIsSorted(tempTests, sortFunc) {
			t.Fatal("incomplete sort")
		}
	})

	t.Run("order by Height desc", func(t *testing.T) {
		t.Log("=== before ===")
		for _, test := range tempTests {
			t.Logf("%f, ", test.Height)
		}
		Order(tempTests).Desc("Height").Exec()
		t.Log("=== after ===")
		for _, test := range tempTests {
			t.Logf("%fs, ", test.Height)
		}

		sortFunc := func(i, j int) bool {
			return tempTests[i].Height > tempTests[j].Height
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
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Asc("ID").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].ID < tempTests[j].ID
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Name_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Asc("Name").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].Name < tempTests[j].Name
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Height_Asc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Asc("Height").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].Height < tempTests[j].Height
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
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Desc("ID").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].ID > tempTests[j].ID
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Name_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Desc("Name").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].Name > tempTests[j].Name
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

func BenchmarkSort_Order_By_Height_Desc(b *testing.B) {
	rand.Seed(time.Now().Unix())

	var tempTests []Test
	for i := 0; i < 10000; i++ {
		tempTests = append(tempTests, Test{
			ID:     rand.Uint64(),
			Name:   randStringRunes(5),
			Age:    uint(rand.Uint32()),
			Height: rand.Float32(),
			Weight: rand.Float32(),
		})
	}

	b.ResetTimer()
	Order(tempTests).Desc("Height").Exec()

	sortFunc := func(i, j int) bool {
		return tempTests[i].Height > tempTests[j].Height
	}
	if !sort.SliceIsSorted(tempTests, sortFunc) {
		b.Fatal("incomplete sort")
	}
}

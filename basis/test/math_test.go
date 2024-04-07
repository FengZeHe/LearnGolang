package test

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		intputA int
		intputB int
		want    int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, tc := range testCases {
		t.Run("Add", func(t *testing.T) {
			if got := Add(tc.intputA, tc.intputB); got != tc.want {
				t.Errorf("got %d ,want %d", got, tc.want)
			}
		})
	}

}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		intputA int
		intputB int
		want    int
	}{
		{1, 1, 0},
		{3, 2, 1},
	}

	for _, tc := range testCases {
		t.Run("Subtract", func(t *testing.T) {
			if got := Subtract(tc.intputA, tc.intputB); got != tc.want {
				t.Errorf("got %d ,want %d", got, tc.want)
			}
		})
	}
}

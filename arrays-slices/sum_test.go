package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 1, 3, 5}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{3, 2})
	want := []int{6, 5}

	if !slices.Equal(got, want) {
		t.Errorf("got %d, want %d given, %v", got, want, [][]int{{1, 2, 3}, {3, 2}})
	}

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}
	t.Run("should calculate tail sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{0, 9, 3})
		want := []int{4, 12}

		checkSums(t, got, want)
	})

	t.Run("should work with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 2, 1})
		want := []int{0, 3}

		checkSums(t, got, want)
	})

}

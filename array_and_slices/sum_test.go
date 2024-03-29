package array_and_slices

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("Expected %d but got %d, given %v", want, got, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("Expected %d but got %d, given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checksums := func(t *testing.T, want, got []int) {
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}

	t.Run("make some without first value", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checksums(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1}, []int{2, 3, 4})
		want := []int{0, 0, 7}

		checksums(t, want, got)
	})
}

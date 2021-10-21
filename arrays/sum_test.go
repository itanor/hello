package arrays

import (
  "testing"
  "reflect"
)

func TestSum(t *testing.T) {

  t.Run("collection of any size", func(t *testing.T) {
    numbers := []int{1, 2, 3}

    got := Sum(numbers)
    want := 6

    if got != want {
      t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
  })

  t.Run("totals for each slice - one slice, one element", func(t *testing.T) {
    first := []int{0}

    got := SumAll(first)
    want := []int{0}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("totals for each slice - one slice, three elements", func(t *testing.T) {
    first := []int{0, 5, 8}
    second := []int{0, 0, 2}
    third := []int{3, 5, -1}

    got := SumAll(first, second, third)
    want := []int{13, 2, 7}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("totals for each slice - two slices, three elements", func(t *testing.T) {
    first := []int{1, 2, 3}
    second := []int{1, 1, 2}

    got := SumAll(first, second)
    want := []int{6, 4}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("sum all tails of each slice - multiple elements in slices", func(t *testing.T) {
    first := []int{1, 2, 3}
    second := []int{1, 1, 2}

    got := SumAllTails(first, second)
    want := []int{5, 3}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("sum all tails of each slice - one element in one slice", func(t *testing.T) {
    first := []int{3}
    second := []int{1, 5}

    got := SumAllTails(first, second)
    want := []int{3, 5}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("sum all tails of each slice - three slices", func(t *testing.T) {
    first := []int{3}
    second := []int{1, 5}
    third := []int{6}

    got := SumAllTails(first, second, third)
    want := []int{3, 5, 6}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })

  t.Run("sum all tails of each slice - three slices, one with zero size", func(t *testing.T) {
    first := []int{}
    second := []int{1, 5, 3}
    third := []int{6}

    got := SumAllTails(first, second, third)
    want := []int{0, 8, 6}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  })
}


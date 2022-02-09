package main

import (
  "testing"
)

func TestSearch(t *testing.T) {
  dictionay := Dictionary{"test": "this is just a test"}

  t.Run("know word", func(t *testing.T) {
    got, _ := dictionay.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })

  t.Run("unknow word", func(t *testing.T) {
    _, got := dictionay.Search("unknown")

    assertError(t, got, ErrNotFound)
  })
}

func assertStrings(t testing.TB, got, want string) {
  t.Helper()

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func assertError(t testing.TB, got, want error) {
  t.Helper()

  if got != want {
    t.Errorf("got error %q want %q", got, want)
  }
}

func TestAdd(t *testing.T) {

  t.Run("new word", func(t *testing.T) {
    dictionay := Dictionary{}
    key := "test"
    value := "this is just a test"
    err := dictionay.Add(key, value)

    assertError(t, err, nil)
    assertValue(t, dictionay, key, value)
  })

  t.Run("existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionay := Dictionary{key: value}
    err := dictionay.Add(key, "new test")

    assertError(t, err, ErrKeyExists)
    assertValue(t, dictionay, key, value)
  })
}

func TestUpdate(t *testing.T) {

  t.Run("existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionay := Dictionary{key: value}
    newValue := "new value"

    err := dictionay.Update(key, newValue)

    assertError(t, err, nil)
    assertValue(t, dictionay, key, newValue)
  })

  t.Run("new key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionay := Dictionary{}

    err := dictionay.Update(key, value)

    assertError(t, err, ErrKeyNotExists)
  })
}

func TestDelete(t *testing.T) {
  key := "test"
  value := "some value"
  dictionay := Dictionary{key: value}

  dictionay.Delete(key)

  _, err := dictionay.Search(key)
  if err != ErrNotFound {
    t.Errorf("Expected %q to be deleted", key)
  }
}

func assertValue(t testing.TB, dictionay Dictionary, key, value string) {
  t.Helper()

  got, err := dictionay.Search(key)
  if err != nil {
    t.Fatal("should find added value", err)
  }
  if value != got {
    t.Errorf("got %q, want %q", got, value)
  }
}

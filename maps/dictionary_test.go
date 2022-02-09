package main

import (
  "testing"
)

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is just a test"}

  t.Run("know word", func(t *testing.T) {
    got, _ := dictionary.Search("test")
    want := "this is just a test"

    assertStrings(t, got, want)
  })

  t.Run("unknow word", func(t *testing.T) {
    _, got := dictionary.Search("unknown")

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
    dictionary := Dictionary{}
    key := "test"
    value := "this is just a test"
    err := dictionary.Add(key, value)

    assertError(t, err, nil)
    assertValue(t, dictionary, key, value)
  })

  t.Run("existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}
    err := dictionary.Add(key, "new test")

    assertError(t, err, ErrKeyExists)
    assertValue(t, dictionary, key, value)
  })
}

func TestUpdate(t *testing.T) {

  t.Run("existing key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{key: value}
    newValue := "new value"

    err := dictionary.Update(key, newValue)

    assertError(t, err, nil)
    assertValue(t, dictionary, key, newValue)
  })

  t.Run("new key", func(t *testing.T) {
    key := "test"
    value := "this is just a test"
    dictionary := Dictionary{}

    err := dictionary.Update(key, value)

    assertError(t, err, ErrKeyNotExists)
  })
}

func TestDelete(t *testing.T) {
  key := "test"
  value := "some value"
  dictionary := Dictionary{key: value}

  dictionary.Delete(key)

  _, err := dictionary.Search(key)
  if err != ErrNotFound {
    t.Errorf("Expected %q to be deleted", key)
  }
}

func assertValue(t testing.TB, dictionary Dictionary, key, value string) {
  t.Helper()

  got, err := dictionary.Search(key)
  if err != nil {
    t.Fatal("should find added value", err)
  }
  if value != got {
    t.Errorf("got %q, want %q", got, value)
  }
}

package main

const (
  ErrNotFound = DictionaryErr("could not find the word")
  ErrKeyExists = DictionaryErr("key exists")
  ErrKeyNotExists = DictionaryErr("key not exists")
)

type DictionaryErr string
type Dictionary map[string]string

func (e DictionaryErr) Error() string {
  return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
  if value, ok := d[key]; ok {
    return value, nil
  }
  return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
  if _, err := d.Search(key); err == nil {
    return ErrKeyExists
  }
  d[key] = value
  return nil
}

func (d Dictionary) Update(key, value string) error {
  if _, err := d.Search(key); err == nil {
    d[key] = value
    return nil
  }
  return ErrKeyNotExists
}

func (d Dictionary) Delete(key string) error {
  if _, err := d.Search(key); err == nil {
    delete(d, key)
    return nil
  }
  return ErrKeyNotExists
}


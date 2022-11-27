package mydict

import "errors"

type Dictioary map[string]string

var errNotFound = errors.New("Not Found")

func (d Dictioary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

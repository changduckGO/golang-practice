package mydict

import "errors"

type Dictioary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictioary) Search(word string) (string, error) {
	def, exists := d[word]
	if exists {
		return def, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictioary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

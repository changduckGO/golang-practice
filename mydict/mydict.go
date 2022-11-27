package mydict

import "errors"

type Dictioary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

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

// Update a word
func (d Dictioary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictioary) Delete(word string) {
	delete(d, word)
}

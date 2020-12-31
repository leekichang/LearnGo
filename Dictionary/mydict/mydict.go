package mydict

import "errors"

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't Update non-existing word")
	errWordExists = errors.New("That word already exusts")
)

//Dictionary tyoe
type Dictionary map[string]string

//Search for word through the dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

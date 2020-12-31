package mydict

import "errors"

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exusts")

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

package mydict

import "errors"

var errNotFound = errors.New("Not Found")

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

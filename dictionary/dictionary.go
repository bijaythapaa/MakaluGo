package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound          = errors.New("could not find the word you are looking for")
	ErrWordExists        = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExists = errors.New("cannot update word because it does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//func Search(dictionary map[string]string, word string) string {
//	return dictionary[word]
//}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		//return "", errors.New("could not find the word you were looking for")
		return "", ErrNotFound
	}
	//return d[word], nil
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	//d[word] = definition
	//return nil

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	//d[word] = definition
	//return nil

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

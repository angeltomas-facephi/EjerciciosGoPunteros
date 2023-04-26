package main

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	book := &Book{
		Name:        "Eloquent JavaScript",
		Release:     2018,
		Language:    "Español, Inglés, Arabe, Persa, Búlgaro",
		Description: "This is a description.",
	}

	expected := `{"name":"Eloquent JavaScript","release":2018,"language":"Español, Inglés, Arabe, Persa, Búlgaro","description":"This is a description."}`

	result, err := serialize(book)

	if err != nil {
		t.Errorf("error	serializing: %s", err)
	}

	if string(result) != expected {
		t.Errorf(`was incorrect, got: "%s"`, expected)
		t.Errorf(`want: "%s"`, string(result))
	}

}

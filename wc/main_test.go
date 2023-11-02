package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("aa bb cc dd\n")
	expected := 4
	result, err := count(input, false, false)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d words\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("aa \n bb \n cc")
	expected := 3
	result, _ := count(input, true, false)
	if result != expected {
		t.Errorf("Expected %d, but got %d lines\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	input := bytes.NewBuffer([]byte{1, 2, 3, 4, 5})
	expected := 5
	result, _ := count(input, false, true)
	if result != expected {
		t.Errorf("Expected %d, but got %d bytes\n", expected, result)
	}
}

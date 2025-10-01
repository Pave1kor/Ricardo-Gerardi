package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("1 2 3")
	exp := 3
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, gor%d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("1 2 3")
	exp := 1
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}

}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("1 2 3")
	exp := 5
	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}

}

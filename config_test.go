package main

import (
	"testing"
)

func TestLoad(t *testing.T) {
	fileName := ""
	config := NewConfig()
	err := config.Load(fileName)
	if err == nil {
		t.Error("Expected nil got", err)
	}
}

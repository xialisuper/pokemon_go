package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Single Word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "Multiple Words",
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Empty Input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "All Uppercase",
			input:    "HELLO THERE",
			expected: []string{"hello", "there"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := cleanInput(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("cleanInput(%s) - expected: %v, got: %v", tt.input, tt.expected, actual)
			}
		})
	}
}

func Test_commands(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
	}{
		{name: "test help command", cmd: "help"},
		{name: "test exit command", cmd: "exit"},
		{name: "test mapf command", cmd: "mapf"},
		{name: "test mapb command", cmd: "mapb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commands()[tt.cmd]
			if got.name == "" {
				t.Errorf("Empty name for command %s", tt.cmd)
			}
			if got.description == "" {
				t.Errorf("Empty description for command %s", tt.cmd)
			}
			if got.callback == nil {
				t.Errorf("Nil callback for command %s", tt.cmd)
			}
		})
	}
}

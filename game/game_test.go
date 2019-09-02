package game

import "testing"

func TestGameIt(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "problem 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GameIt()
		})
	}
}

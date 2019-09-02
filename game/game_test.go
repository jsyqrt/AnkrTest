package game

import "testing"

func TestGameIt(t *testing.T) {
	type args struct {
		game Game
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "problem 2",
			args: args{
				game: &GameA{
					NewEmptyCheckerBoard(),
				},
			},
		},
		{
			name: "problem 3",
			args: args{
				game: &GameB{
					NewEmptyCheckerBoard(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GameIt(tt.args.game)
		})
	}
}

package game

import "testing"

func TestCheckerBoard_Print(t *testing.T) {
	type fields struct {
		board [RowSize][ColSize]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "case1:empty",
			fields: fields{
				board: NewEmptyCheckerBoard().board,
			},
		},
		{
			name: "case2",
			fields: fields{
				board: func() [RowSize][ColSize]int {
					m := NewEmptyCheckerBoard().board
					m[0][3], m[3][2] = 1, 1
					m[2][3], m[3][5] = 2, 2
					return m
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &CheckerBoard{
				board: tt.fields.board,
			}
			b.Print()
		})
	}
}

package test

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
		x float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_1",
			args: args{
				a: -8,
				x: -5,
				b: -2,
				c: 7,
			},
		},
		{
			name: "Test_2",
			args: args{
				a: 8,
				x: 2,
				b: 9,
				c: -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.args.a, tt.args.x, tt.args.b, tt.args.c)
		})
	}
}

package testrun

import "testing"

func TestTestRun_RunInt(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := &TestRun{}
			gotResult, err := te.RunInt(tt.args.in...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("RunInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

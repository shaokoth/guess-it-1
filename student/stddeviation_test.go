package stats

import "testing"

func TestStandardDev(t *testing.T) {
	type args struct {
		input []float64
		mean  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test2",
			args: args{input: []float64{1, 2, 3, 4, 5}},
			want: 3.3166247903554,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDev(tt.args.input, tt.args.mean); got != tt.want {
				t.Errorf("StandardDev() = %v, want %v", got, tt.want)
			}
		})
	}
}

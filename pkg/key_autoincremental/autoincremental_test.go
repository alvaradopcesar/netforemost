package key_autoincremental

import "testing"

func Test_incremental_Next(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test 01",
			want: 1,
		},
		{
			name: "test 02",
			want: 2,
		},
		{
			name: "test 03",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := incremental{}
			if got := i.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

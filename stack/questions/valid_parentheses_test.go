package questions

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "one char string",
			s:    "{",
			want: false,
		},
		{
			name: "test case 1",
			s:    "()",
			want: true,
		},
		{
			name: "test case 2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "test case 3",
			s:    "(]",
			want: false,
		},
		{
			name: "test case 4",
			s:    "([])",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

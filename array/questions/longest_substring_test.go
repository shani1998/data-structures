package questions

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Test case 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Test case 4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test case 5",
			args: args{
				s: " ",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
			if got := lengthOfLongestSubstring2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring2() = %v, want %v", got, tt.want)
			}
		})
	}
}

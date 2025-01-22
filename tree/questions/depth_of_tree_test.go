package questions

import (
	"reflect"
	"testing"

	"github.com/shani1998/data-structures/tree"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *tree.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root: &tree.Node{
					Val: 3,
				},
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				root: &tree.Node{
					Val: 3,
					Left: &tree.Node{
						Val: 9,
					},
					Right: &tree.Node{
						Val: 20,
						Left: &tree.Node{
							Val: 15,
						},
						Right: &tree.Node{
							Val: 7,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthPath(t *testing.T) {
	type args struct {
		root *tree.Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Single node tree",
			args: args{
				root: &tree.Node{
					Val: 3,
				},
			},
			want: []int{3},
		},
		{
			name: "Two-level tree",
			args: args{
				root: &tree.Node{
					Val:  3,
					Left: &tree.Node{Val: 9},
				},
			},
			want: []int{3, 9},
		},
		{
			name: "Three-level tree with both left and right subtrees",
			args: args{
				root: &tree.Node{
					Val: 3,
					Left: &tree.Node{
						Val:  9,
						Left: &tree.Node{Val: 4},
					},
					Right: &tree.Node{
						Val: 20,
						Left: &tree.Node{
							Val: 15,
						},
						Right: &tree.Node{
							Val: 7,
						},
					},
				},
			},
			want: []int{3, 20, 7}, // Longest path is to the right subtree
		},
		{
			name: "Skewed left tree",
			args: args{
				root: &tree.Node{
					Val: 3,
					Left: &tree.Node{
						Val:  9,
						Left: &tree.Node{Val: 4},
					},
				},
			},
			want: []int{3, 9, 4}, // Longest path is the left branch
		},
		{
			name: "Skewed right tree",
			args: args{
				root: &tree.Node{
					Val: 3,
					Right: &tree.Node{
						Val: 20,
						Right: &tree.Node{
							Val: 7,
						},
					},
				},
			},
			want: []int{3, 20, 7}, // Longest path is the right branch
		},
		{
			name: "Empty tree",
			args: args{
				root: nil,
			},
			want: []int{}, // No nodes in the tree
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthPath(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

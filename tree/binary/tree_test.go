package binary

import (
	"reflect"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		objs []Comparor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *node
		wantErr bool
	}{
		{
			name: "Insert Smaller Node",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
				},
			},
			args: args{
				objs: []Comparor{
					&mockIntCompare{
						data: 30,
					},
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 50,
				},
				left: &node{
					data: &mockIntCompare{
						data: 30,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Insert Larger Node",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
				},
			},
			args: args{
				objs: []Comparor{
					&mockIntCompare{
						data: 70,
					},
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 50,
				},
				right: &node{
					data: &mockIntCompare{
						data: 70,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Insert Tree",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
				},
			},
			args: args{
				objs: []Comparor{
					&mockIntCompare{
						data: 70,
					},
					&mockIntCompare{
						data: 30,
					},
					&mockIntCompare{
						data: 40,
					},
					&mockIntCompare{
						data: 60,
					},
					&mockIntCompare{
						data: 20,
					},
					&mockIntCompare{
						data: 80,
					},
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 50,
				},
				left: &node{
					data: &mockIntCompare{
						data: 30,
					},
					left: &node{
						data: &mockIntCompare{
							data: 20,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 40,
						},
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 70,
					},
					left: &node{
						data: &mockIntCompare{
							data: 60,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 80,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}

			for _, obj := range tt.args.objs {
				if err := tree.Insert(obj); (err != nil) != tt.wantErr {
					t.Errorf("Tree.Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			if reflect.DeepEqual(tree.root, tt.want) == false {
				t.Errorf("Tree.Insert() got = %v, want %v", tree.root, tt.want)
			}
		})
	}
}

func TestTree_get(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		obj Comparor
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantParent *node
		wantChild  *node
	}{
		{
			name: "Left Side",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 20,
				},
			},
			wantParent: &node{
				data: &mockIntCompare{
					data: 30,
				},
				left: &node{
					data: &mockIntCompare{
						data: 20,
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 40,
					},
				},
			},
			wantChild: &node{
				data: &mockIntCompare{
					data: 20,
				},
			},
		},
		{
			name: "Right Side",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 80,
				},
			},
			wantParent: &node{
				data: &mockIntCompare{
					data: 70,
				},
				left: &node{
					data: &mockIntCompare{
						data: 60,
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 80,
					},
				},
			},
			wantChild: &node{
				data: &mockIntCompare{
					data: 80,
				},
			},
		},
		{
			name: "No Child",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 10,
				},
			},
			wantParent: &node{
				data: &mockIntCompare{
					data: 20,
				},
			},
			wantChild: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			gotParent, gotChild := tree.get(tree.root, tt.args.obj)
			if !reflect.DeepEqual(gotParent, tt.wantParent) {
				t.Errorf("Tree.get() gotParent = %v, want %v", gotParent, tt.wantParent)
			}
			if !reflect.DeepEqual(gotChild, tt.wantChild) {
				t.Errorf("Tree.get() gotChild = %v, want %v", gotChild, tt.wantChild)
			}
		})
	}
}

func TestTree_Has(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		obj Comparor
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Has",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 70,
				},
			},
			want: true,
		},
		{
			name: "Does not have",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 12,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.Has(tt.args.obj); got != tt.want {
				t.Errorf("Tree.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_depth(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n *node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "depth",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.depth(tt.args.n); got != tt.want {
				t.Errorf("Tree.depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Depth(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		path Path
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Root",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				path: Root,
			},
			want: 3,
		},
		{
			name: "Left",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
					},
				},
			},
			args: args{
				path: Left,
			},
			want: 2,
		},
		{
			name: "Left",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
					},
				},
			},
			args: args{
				path: Right,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.Depth(tt.args.path); got != tt.want {
				t.Errorf("Tree.Depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		root Comparor
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "new",
			args: args{
				root: &mockIntCompare{
					data: 50,
				},
			},
			want: &Tree{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_inorder(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n    *node
		objs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			name:   "In order",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
				objs: []interface{}{},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 50,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 80,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.inorder(tt.args.n, tt.args.objs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.inorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Inorder(t *testing.T) {
	type fields struct {
		root *node
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "In order",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 50,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 80,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.Inorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Inorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_preorder(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n    *node
		objs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			name:   "Pre-order",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
				objs: []interface{}{},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 50,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 80,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.preorder(tt.args.n, tt.args.objs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Preorder(t *testing.T) {
	type fields struct {
		root *node
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "Pre order",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 50,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 80,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.Preorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_postorder(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n    *node
		objs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			name:   "Post-order",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
				objs: []interface{}{},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 80,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.postorder(tt.args.n, tt.args.objs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Postorder(t *testing.T) {
	type fields struct {
		root *node
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "Post order",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			want: []interface{}{
				&mockIntCompare{
					data: 20,
				},
				&mockIntCompare{
					data: 40,
				},
				&mockIntCompare{
					data: 30,
				},
				&mockIntCompare{
					data: 60,
				},
				&mockIntCompare{
					data: 80,
				},
				&mockIntCompare{
					data: 70,
				},
				&mockIntCompare{
					data: 50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.Postorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_minNode(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		n *node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			name:   "Min from root",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 20,
				},
			},
		},
		{
			name:   "Min from right",
			fields: fields{},
			args: args{
				n: &node{
					data: &mockIntCompare{
						data: 70,
					},
					left: &node{
						data: &mockIntCompare{
							data: 60,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 80,
						},
					},
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 60,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.minNode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.minNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Delete(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		obj Comparor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *node
		wantErr bool
	}{
		{
			name: "Delete Leaf",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 20,
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 50,
				},
				left: &node{
					data: &mockIntCompare{
						data: 30,
					},
					right: &node{
						data: &mockIntCompare{
							data: 40,
						},
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 70,
					},
					left: &node{
						data: &mockIntCompare{
							data: 60,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 80,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Delete has only one child",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 30,
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 50,
				},
				left: &node{
					data: &mockIntCompare{
						data: 40,
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 70,
					},
					left: &node{
						data: &mockIntCompare{
							data: 60,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 80,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Delete has two children",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 40,
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				obj: &mockIntCompare{
					data: 50,
				},
			},
			want: &node{
				data: &mockIntCompare{
					data: 60,
				},
				left: &node{
					data: &mockIntCompare{
						data: 40,
					},
				},
				right: &node{
					data: &mockIntCompare{
						data: 70,
					},
					right: &node{
						data: &mockIntCompare{
							data: 80,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if err := tree.Delete(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Tree.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(tree.root, tt.want) == false {
				t.Errorf("Tree.Delete() got = %v, want %v", tree.root, tt.want)
			}
		})
	}
}

func TestTree_subArray(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		main []interface{}
		sub  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Has Same",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: true,
		},
		{
			name:   "Has",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
					&mockIntCompare{
						data: 4,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: true,
		},
		{
			name:   "Nope",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 4,
					},
					&mockIntCompare{
						data: 8,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: false,
		},
		{
			name:   "Smaller",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.subArray(tt.args.main, tt.args.sub); got != tt.want {
				t.Errorf("Tree.subArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_in(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		main []interface{}
		sub  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "In Same",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: true,
		},
		{
			name:   "In",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 6,
					},
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: true,
		},
		{
			name:   "Nope",
			fields: fields{},
			args: args{
				main: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 5,
					},
					&mockIntCompare{
						data: 8,
					},
				},
				sub: []interface{}{
					&mockIntCompare{
						data: 1,
					},
					&mockIntCompare{
						data: 8,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.in(tt.args.main, tt.args.sub); got != tt.want {
				t.Errorf("Tree.in() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_SubTree(t *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		sub *Tree
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Has Sub Tree",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				sub: &Tree{
					root: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "No Sub Tree",
			fields: fields{
				root: &node{
					data: &mockIntCompare{
						data: 50,
					},
					left: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 40,
							},
						},
					},
					right: &node{
						data: &mockIntCompare{
							data: 70,
						},
						left: &node{
							data: &mockIntCompare{
								data: 60,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 80,
							},
						},
					},
				},
			},
			args: args{
				sub: &Tree{
					root: &node{
						data: &mockIntCompare{
							data: 30,
						},
						left: &node{
							data: &mockIntCompare{
								data: 20,
							},
						},
						right: &node{
							data: &mockIntCompare{
								data: 45,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				root: tt.fields.root,
			}
			if got := tree.SubTree(tt.args.sub); got != tt.want {
				t.Errorf("Tree.SubTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

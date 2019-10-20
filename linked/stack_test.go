package linked

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Stack
		wantErr bool
	}{
		{
			name: "Push Empty",
			fields: fields{
				list: &List{},
			},
			args: args{
				data: 5,
			},
			want: &Stack{
				list: &List{
					start: &node{
						data: 5,
					},
					end: &node{
						data: 5,
					},
					size: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Push",
			fields: fields{
				list: &List{
					start: &node{
						data: 5,
					},
					size: 1,
				},
			},
			args: args{
				data: 34,
			},
			want: &Stack{
				list: &List{
					start: &node{
						data: 34,
						next: &node{
							data: 5,
						},
					},
					end: &node{
						data: 5,
					},
					size: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			s.list.end = s.list.start
			if s.list.end != nil {
				for s.list.end.next != nil {
					s.list.end = s.list.end.next
				}
			}
			if err := s.Push(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(s, tt.want) == false {
				t.Errorf("Stack.Push() got = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name      string
		fields    fields
		want      interface{}
		wantStack *Stack
		wantErr   bool
	}{
		{
			name: "Pop",
			fields: fields{
				list: &List{
					start: &node{
						data: 34,
						next: &node{
							data: 5,
						},
					},
					size: 2,
				},
			},
			want: 34,
			wantStack: &Stack{
				list: &List{
					start: &node{
						data: 5,
					},
					end: &node{
						data: 5,
					},
					size: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Pop Last",
			fields: fields{
				list: &List{
					start: &node{
						data: 5,
					},
					size: 1,
				},
			},
			want: 5,
			wantStack: &Stack{
				list: &List{},
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				list: &List{},
			},
			want: nil,
			wantStack: &Stack{
				list: &List{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			s.list.end = s.list.start
			if s.list.end != nil {
				for s.list.end.next != nil {
					s.list.end = s.list.end.next
				}
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() data = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual(s, tt.wantStack) == false {
				t.Errorf("Stack.Pop() got = %v, want %v", s, tt.wantStack)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name      string
		fields    fields
		want      interface{}
		wantStack *Stack
		wantErr   bool
	}{
		{
			name: "Pop",
			fields: fields{
				list: &List{
					start: &node{
						data: 34,
						next: &node{
							data: 5,
						},
					},
					size: 2,
				},
			},
			want: 34,
			wantStack: &Stack{
				list: &List{
					start: &node{
						data: 34,
						next: &node{
							data: 5,
						},
					},
					size: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			got, err := s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() data = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual(s, tt.wantStack) == false {
				t.Errorf("Stack.Pop() got = %v, want %v", s, tt.wantStack)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "Size",
			fields: fields{
				list: &List{
					size: 4,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				list: tt.fields.list,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

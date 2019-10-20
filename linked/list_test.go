package linked

import (
	"errors"
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	type fields struct {
		start *node
		size  int
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *List
		wantErr bool
	}{
		{
			name:   "Append from the begining",
			fields: fields{},
			args: args{
				data: "starting",
			},
			want: &List{
				start: &node{
					data: "starting",
				},
				size: 1,
			},
			wantErr: false,
		},
		{
			name: "Append from the end",
			fields: fields{
				start: &node{
					data: "starting",
				},
				size: 1,
			},
			args: args{
				data: "ending",
			},
			want: &List{
				start: &node{
					data: "starting",
					next: &node{
						data: "ending",
					},
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name: "Append from the end too",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
					},
				},
				size: 1,
			},
			args: args{
				data: "ending",
			},
			want: &List{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
						next: &node{
							data: "ending",
						},
					},
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name:    "Error",
			fields:  fields{},
			args:    args{},
			want:    &List{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				size:  tt.fields.size,
			}
			if err := l.Append(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("List.Append() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(l, tt.want) == false {
				t.Errorf("List.Append() got = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestList_Delete(t *testing.T) {
	type fields struct {
		start *node
		size  int
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *List
		wantErr bool
	}{
		{
			name: "Delete from start",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 0,
			},
			want: &List{
				start: &node{
					data: 4,
					next: &node{
						data: 9,
					},
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name: "Delete from end",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 2,
			},
			want: &List{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
					},
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name: "Delete from middle",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 1,
			},
			want: &List{
				start: &node{
					data: 1,
					next: &node{
						data: 9,
					},
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 5,
			},
			want: &List{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				size:  tt.fields.size,
			}
			if err := l.Delete(tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("List.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(l, tt.want) == false {
				t.Errorf("List.Delete() got = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestList_Size(t *testing.T) {
	type fields struct {
		start *node
		size  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "three nodes",
			fields: fields{
				size: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				size:  tt.fields.size,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("List.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_ForEach(t *testing.T) {
	type fields struct {
		start *node
		size  int
	}
	type args struct {
		f func(data interface{}, idx int) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Iterate",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				f: func(data interface{}, idx int) error {
					n, ok := data.(int)
					if ok == false {
						return errors.New("not a number")
					}
					list := []int{1, 4, 9}
					if list[idx] != n {
						return errors.New("not equal")
					}
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				start: &node{
					data: 1,
					next: &node{
						data: 4,
						next: &node{
							data: 9,
						},
					},
				},
				size: 3,
			},
			args: args{
				f: func(data interface{}, idx int) error {
					return errors.New("something")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				size:  tt.fields.size,
			}
			if err := l.ForEach(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("List.ForEach() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

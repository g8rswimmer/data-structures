package linked

import (
	"errors"
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	type fields struct {
		start *node
		size  uint
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
				end: &node{
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
				end: &node{
					data: "ending",
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
				end: &node{
					data: "ending",
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
				end:   tt.fields.start,
				size:  tt.fields.size,
			}
			if l.end != nil {
				for l.end.next != nil {
					l.end = l.end.next
				}
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
		size  uint
	}
	type args struct {
		idx uint
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
				end: &node{
					data: 9,
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
				end: &node{
					data: 4,
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
				end: &node{
					data: 9,
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
				end: &node{
					data: 9,
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
				end:   tt.fields.start,
				size:  tt.fields.size,
			}
			if l.end != nil {
				for l.end.next != nil {
					l.end = l.end.next
				}
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
		size  uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
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
		size  uint
	}
	type args struct {
		f func(data interface{}, idx uint) error
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
				f: func(data interface{}, idx uint) error {
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
				f: func(data interface{}, idx uint) error {
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

func TestList_Insert(t *testing.T) {
	type fields struct {
		start *node
		size  uint
	}
	type args struct {
		data interface{}
		idx  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *List
		wantErr bool
	}{
		{
			name:   "Insert from the begining",
			fields: fields{},
			args: args{
				data: "starting",
				idx:  0,
			},
			want: &List{
				start: &node{
					data: "starting",
				},
				end: &node{
					data: "starting",
				},
				size: 1,
			},
			wantErr: false,
		},
		{
			name: "Insert from the end",
			fields: fields{
				start: &node{
					data: "starting",
				},
				size: 1,
			},
			args: args{
				data: "ending",
				idx:  1,
			},
			want: &List{
				start: &node{
					data: "starting",
					next: &node{
						data: "ending",
					},
				},
				end: &node{
					data: "ending",
				},
				size: 2,
			},
			wantErr: false,
		},
		{
			name: "Insert from the end",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "ending",
					},
				},
				size: 2,
			},
			args: args{
				data: "middle",
				idx:  1,
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
				end: &node{
					data: "ending",
				},
				size: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				end:   tt.fields.start,
				size:  tt.fields.size,
			}
			if l.end != nil {
				for l.end.next != nil {
					l.end = l.end.next
				}
			}
			if err := l.Insert(tt.args.data, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("List.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(l, tt.want) == false {
				t.Errorf("List.Insert() got = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestList_Retrieve(t *testing.T) {
	type fields struct {
		start *node
		size  uint
	}
	type args struct {
		idx uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Retrieve from the start",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
						next: &node{
							data: "ending",
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 0,
			},
			want:    "starting",
			wantErr: false,
		},
		{
			name: "Retrieve from the end",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
						next: &node{
							data: "ending",
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 2,
			},
			want:    "ending",
			wantErr: false,
		},
		{
			name: "Retrieve from the middle",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
						next: &node{
							data: "ending",
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 1,
			},
			want:    "middle",
			wantErr: false,
		},
		{
			name: "Error",
			fields: fields{
				start: &node{
					data: "starting",
					next: &node{
						data: "middle",
						next: &node{
							data: "ending",
						},
					},
				},
				size: 3,
			},
			args: args{
				idx: 5,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				start: tt.fields.start,
				end:   tt.fields.start,
				size:  tt.fields.size,
			}
			if l.end != nil {
				for l.end.next != nil {
					l.end = l.end.next
				}
			}
			got, err := l.Retrieve(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

package linked

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
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
		want    *Queue
		wantErr bool
	}{
		{
			name: "Enqueue empty",
			fields: fields{
				list: NewList(),
			},
			args: args{
				data: 4,
			},
			want: &Queue{
				list: &List{
					start: &node{
						data: 4,
					},
					end: &node{
						data: 4,
					},
					size: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "Enqueue",
			fields: fields{
				list: &List{
					start: &node{
						data: 4,
					},
					size: 1,
				},
			},
			args: args{
				data: 78,
			},
			want: &Queue{
				list: &List{
					start: &node{
						data: 4,
						next: &node{
							data: 78,
						},
					},
					end: &node{
						data: 78,
					},
					size: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			q.list.end = q.list.start
			if q.list.end != nil {
				for q.list.end.next != nil {
					q.list.end = q.list.end.next
				}
			}
			if err := q.Enqueue(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Queue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(q, tt.want) == false {
				t.Errorf("Queue.Enqueue() got = %v, want %v", q, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name      string
		fields    fields
		want      interface{}
		wantQueue *Queue
		wantErr   bool
	}{
		{
			name: "Dequeue",
			fields: fields{
				list: &List{
					start: &node{
						data: 4,
						next: &node{
							data: 78,
						},
					},
					size: 2,
				},
			},
			want: 4,
			wantQueue: &Queue{
				list: &List{
					start: &node{
						data: 78,
					},
					end: &node{
						data: 78,
					},
					size: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			q.list.end = q.list.start
			if q.list.end != nil {
				for q.list.end.next != nil {
					q.list.end = q.list.end.next
				}
			}
			got, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual(q, tt.wantQueue) == false {
				t.Errorf("Queue.Enqueue() got = %v, want %v", q, tt.wantQueue)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
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
					start: &node{
						data: 4,
						next: &node{
							data: 78,
						},
					},
					size: 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				list: tt.fields.list,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

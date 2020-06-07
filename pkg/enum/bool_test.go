package enum

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want BoolEnum
	}{
		{
			name: "success",
			args: args{v: 1},
			want: BoolEnum{v: 1},
		},
		{
			name: "success",
			args: args{v: 0},
			want: BoolEnum{v: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolEnum_ISFalse(t *testing.T) {
	type fields struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "true",
			fields: fields{v: 1},
			want:   false,
		},
		{
			name:   "true",
			fields: fields{v: 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoolEnum{
				v: tt.fields.v,
			}
			if got := b.ISFalse(); got != tt.want {
				t.Errorf("ISFalse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolEnum_IsTrue(t *testing.T) {
	type fields struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "true",
			fields: fields{v: 1},
			want:   true,
		},
		{
			name:   "true",
			fields: fields{v: 0},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoolEnum{
				v: tt.fields.v,
			}
			if got := b.IsTrue(); got != tt.want {
				t.Errorf("IsTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

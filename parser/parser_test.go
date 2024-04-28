package parser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseFields(t *testing.T) {

	mockFields := strings.Fields("1 1 1 1 1 /usr/bin/find")

	type args struct {
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "invalid arguments, return error",
			args:    args{fields: []string{}},
			wantErr: true,
		},
		{
			name:    "valid arguments, return valid string",
			args:    args{fields: mockFields},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseFields(tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_expandField(t *testing.T) {
	type args struct {
		field  string
		minVal int
		maxVal int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "when field have *",
			args: args{field: "*", minVal: 0, maxVal: 6},
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "when field have /",
			args: args{field: "/2", minVal: 1, maxVal: 12},
			want: []int{1, 3, 5, 7, 9, 11},
		},
		{
			name: "when field have -",
			args: args{field: "1-5", minVal: 1, maxVal: 12},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandField(tt.args.field, tt.args.minVal, tt.args.maxVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandField() = %v, want %v", got, tt.want)
			}
		})
	}
}

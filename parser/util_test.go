package parser

import (
	"testing"
)

func Test_intSliceToString(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid argument, return valid string",
			args: args{slice: []int{1, 2, 3, 4, 5}},
			want: "1 2 3 4 5",
		},
		{
			name: "empty arguments, return empty string",
			args: args{slice: []int{}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intSliceToString(tt.args.slice); got != tt.want {
				t.Errorf("intSliceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cronToString(t *testing.T) {

	mockOutput := "minute\t\t" + "1 2" + "\n"
	mockOutput += "hour\t\t" + "1 2" + "\n"
	mockOutput += "day of month\t" + "1 2" + "\n"
	mockOutput += "month\t\t" + "1 2" + "\n"
	mockOutput += "day of week\t" + "1 2" + "\n"
	mockOutput += "command\t\t" + "/urs/bin/find"

	type args struct {
		s schedule
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid arguments",
			args: args{s: schedule{minutes: []int{1, 2},
				hours: []int{1, 2}, daysOfMonth: []int{1, 2}, months: []int{1, 2}, daysOfWeek: []int{1, 2}, command: "/urs/bin/find"}},
			want: mockOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cronToString(tt.args.s); got != tt.want {
				t.Errorf("cronToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateString(t *testing.T) {

	mockInput := []string{"*", "*", "*", "*", "*", "/usr/bin/find"}

	type args struct {
		fields []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case for valid string expression, return nil",
			args:    args{fields: mockInput},
			wantErr: false,
		},
		{
			name:    "case for invalid string expression for month, return error",
			args:    args{fields: []string{"*", "*", "*", "0/12", "7", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for day of month, return error",
			args:    args{fields: []string{"*", "*", "*/0", "*", "7", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for day of week, return error",
			args:    args{fields: []string{"*", "*", "*", "*", "1-7", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for minute, return error",
			args:    args{fields: []string{"60", "*", "*", "*", "*", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for hour when value not integer, return error",
			args:    args{fields: []string{"*", "h", "*", "*", "1-7", "/usr/bin/find"}},
			wantErr: true,
		},

		{
			name:    "case for invalid string expression for range, return error",
			args:    args{fields: []string{"-1", "h", "*", "*", "*", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for range with non-integer value, return error",
			args:    args{fields: []string{"1-a", "h", "*", "*", "*", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for / with non-integer value, return error",
			args:    args{fields: []string{"a/1", "h", "*", "*", "*", "/usr/bin/find"}},
			wantErr: true,
		},
		{
			name:    "case for invalid string expression for / , return error",
			args:    args{fields: []string{"1/", "h", "*", "*", "*", "/usr/bin/find"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateString(tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("validateString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

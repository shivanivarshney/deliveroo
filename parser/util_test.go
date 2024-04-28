package parser

import (
	"strings"
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

	mockFields := strings.Fields("* * * * * /usr/bin/find")

	type args struct {
		expression []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "invalid arguments",
			args:    args{expression: []string{}},
			wantErr: true,
		},
		{
			name:    "valid arguments",
			args:    args{expression: mockFields},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateString(tt.args.expression); (err != nil) != tt.wantErr {
				t.Errorf("validateString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

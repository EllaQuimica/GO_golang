package main

import "testing"

func Test_calculate_operate(t *testing.T) {
	type args struct {
		input     string
		operation string
	}
	var tests []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calculate{}
			got, err := c.operate(tt.args.input, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("operate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("operate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate_parseString(t *testing.T) {
	type args struct {
		operator string
	}
	var tests []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := calculate{}
			got, err := ca.parseString(tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processResult(t *testing.T) {
	type args struct {
		input    string
		operator string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_readInput(t *testing.T) {
	var tests []struct {
		name string
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(); got != tt.want {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
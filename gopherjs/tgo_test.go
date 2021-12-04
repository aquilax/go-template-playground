package main

import "testing"

func Test_render(t *testing.T) {
	type args struct {
		t    string
		data string
	}
	tests := []struct {
		name string
		args args
		want string
		err  string
	}{
		{
			"simple test",
			args{
				"{{- .id -}}",
				`{"id": "test"}`,
			},
			"test",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := render(tt.args.t, tt.args.data)
			if err != tt.err {
				t.Errorf("render() err = %v, want %v", err, tt.err)
			}
			if got != tt.want {
				t.Errorf("render() got = %v, want %v", got, tt.want)
			}

		})
	}
}

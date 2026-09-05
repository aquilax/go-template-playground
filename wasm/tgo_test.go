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

func Test_compress_decompress(t *testing.T) {
	t.Run("test compress/decompress", func(t *testing.T) {
		input := `{"test": "data"}`
		compressed := "qlYqSS0uUbJSUEpJLElUqgUEAAD__w"
		comp, err := compress(input)
		if err != "" {
			t.Errorf("got compress error = %s", err)
		}
		if comp != compressed {
			t.Errorf("got compressed = %s, want %s", comp, compressed)
		}
		result, err := decompress(comp)
		if err != "" {
			t.Errorf("got decompress error = %s", err)
		}
		if input != result {
			t.Errorf("got = %s, want %s", result, input)
		}
	})
}

package parserv2

import (
	"testing"
)

func TestParser(t *testing.T) {
	//testcases := []struct {
	//	input   string
	//	wantErr string
	//}{{
	//	input:   "4085551212",
	//	wantErr: "",
	//}, {
	//	input:   "408-555-1212",
	//	wantErr: "",
	//}, {
	//	input:   "408-5551212",
	//	wantErr: "syntax error",
	//}}
	testcases := []struct {
		input   string
		wantErr string
	}{{
		input:   "4085551212",
		wantErr: "",
	}}
	for _, tc := range testcases {
		_, err := Parse([]byte(tc.input))
		var gotErr string
		if err != nil {
			gotErr = err.Error()
		}
		if gotErr != tc.wantErr {
			t.Errorf("%s err: %v, want %v", tc.input, gotErr, tc.wantErr)
		}
	}
}

package main

import "testing"

func TestGreeting_English(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}
	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello World",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
	}

	for _, tc := range tests {
		got := greeting(tc.lang)
		if got != tc.want {
			t.Errorf("Hello world got %s", got)
		}
	}
}

// Copyright (c) 2017 AnUnnamedProject
// Distributed under the MIT software license, see the accompanying
// file LICENSE or http://www.opensource.org/licenses/mit-license.php.

package i18n

import "testing"

func TestTranslate(t *testing.T) {
	if err := Load("translations"); err != nil {
		t.Fatal(err)
	}

	t.Parallel()

	var tests = []struct {
		value    string
		language string
		expected string
		params   interface{}
	}{
		{"Hello World", "", "Hello World", ""},
		{"Hello World", "it", "Ciao Mondo", ""},
		{"Hello World", "es", "Hola Mundo", ""},
		{"Hello Gopher", "", "Hello Gopher", ""},
	}
	for _, test := range tests {
		i18n := New(test.language)
		res := i18n.Print(test.value, test.params)
		if res != test.expected {
			t.Errorf("%s, expected %v, got %v", test.value, test.expected, res)
		}
	}
}

func TestPlural(t *testing.T) {
	if err := Load("translations"); err != nil {
		t.Fatal(err)
	}

	t.Parallel()

	var tests = []struct {
		value      int
		language   string
		zero       string
		one        string
		many       string
		expected   string
		customLang string
	}{
		{-1, "", "no records found.", "one record found.", "%d records found.", "no records found.", ""},
		{0, "it", "no records found.", "one record found.", "%d records found.", "no records found.", "it"},
		{1, "es", "no records found.", "one record found.", "%d records found.", "one record found.", "es"},
		{2, "en", "no records found.", "one record found.", "%d records found.", "2 records found.", "en"},
	}
	for _, test := range tests {
		i18n := New(test.language)
		res := i18n.Plural(test.value, test.zero, test.one, test.many, test.customLang)
		if res != test.expected {
			t.Errorf("%d, expected %v, got %v", test.value, test.expected, res)
		}
	}
}

func BenchmarkPrint(b *testing.B) {
	b.ReportAllocs()

	if err := Load("translations"); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		i18n := New("it")
		i18n.Print("Hello world")
	}
}

func BenchmarkPrintf(b *testing.B) {
	b.ReportAllocs()

	if err := Load("translations"); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		i18n := New("it")
		i18n.Print("Hello world %d", 20)
	}
}

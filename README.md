i18n
====

JSON based transations written in Golang.

## Installation
---------------

```
go get github.com/AnUnnamedProject/i18n
```

## Benchmark
------------

Tested on i7-5820k@4.2GHz (6 core / 12 threads)

```
BenchmarkPrint-12     	30000000	        54.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrintf-12    	10000000	       229 ns/op	      24 B/op	       2 allocs/op
```

## Usage
--------

Create your translation file as a JSON object and put in your project folder (i.e. i18n).

```json
{
	"Hello world": "Hello world"
}
```

the scan the directory for json translation files and prepare the internal translation map:

```go
i18n.Load("i18n")
```

If you need to check for missing languages or translations, enable the Debug

```go
i18n.Debug(true)
```

Warnings are printed directly to stdout via log package.

To instantiate and use the library in your code:

```
// Get a new pointer
tr := i18n.New("es")

// Change language
tr.SetLang("it")

// Get the current language
tr.GetLang()

// Print a translation
str1 := tr.Print("Hello world")

// Printf a translation
str2 := tr.Print("Hello %s", "name")

// Plural
str3 := tr.Plural(10, "No records.", "One Record", "%d Records.")
```

## TODO
-------

- App to extract strings from go and html source files.

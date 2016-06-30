UnsafeJSON
==========

__Use at your own risk.__

A fork of Go's encoding/json package that doesn't escape the marshaled JSON.  Go's stdlib JSON package does escaping for safety reasons and should be preferred over this package.  Only use this if it is safe to not escape the JSON.  Improper usage of this package may lead to security issues and other problems.

This package also contains some convenience funcs and a convenience struct, which are described below.

All code contained in this package, with the exception of the modifications necessary to enable the changes listed below, along with the updating of affected tests, were originally written by The Go Authors. 

Usage: Same as the Go encode/json package. Use when you don't want the marshalled JSON to be HTMLEscaped. Please note that the JSON output from this package is not safe for use in browsers or use within the HTML `<script>` tags. For those scenarios, Go's enconding/json package should be used as it is designed for those use cases.

Changes:
* Elided code responsible for the HTML escaping of `<`, `>`, `&` and associated tests.
* Elided HTMLEscape() and related tests.
* Added MarshalIndentToString() function, which wraps MarshalIndent and returns a string.* 

## Elided HTML Escape functionality
Main purpose, at time of creation, was to provide a marshaller that doesn't HTMLEscape the marshalled bytes. While this is useful for HTML and browser usage, this is neither helpful nor useful in some situations. One example is the creation of [Packer](www.packer.io) templates, which may have inlined bash commands. HTMLEscaping the bash commands is not desirable, as all of the original encoding values must be preserved.

The strings which are no longer HTMLEscaped are `<`, `>`, `&`

## Added MarshalIndentToString(f(v interface{}, prefix, indent string) string {}
MarshalIndentToString provides a convenient way of converting something to JSON and then returning a *string* version of it. Since MarshallIndentToString only has a single return value, which is a *string*, it is easy to use whenever you want to see a JSON version of a struct in a formatted string version.  This is helpful in debugging because it can expose some information that can be elided when being printed out, e.g. nil being displayed as nulls, display of quotees, etc. Adding the Indent to the string output makes complex structures more parsable by humans.

    fmt.Println(json.MarshalIndentToString(someInterface, "", "        "))
    
The above would print a JSON version of `someInterface{}` as a formatted string.

MarshalIndentToString wraps a call to MarshalIndent. If an error occurs, the error information is discarded an an empty string, `""`, is returned. If you need the error information, call MarshalIndent() instead. Otherwise, the *[]byte* is converted to a *string* and returned.

## Added MarshalToString(f(v interface{}) string {}
MarshallToString provides a convenient way of converting something to JSON and then returning a *string* version of it. Since MarshallToString only has a single return value, which is a *string*, it is easy to use whenever you want to see a JSON version of a struct. This is helpful in debugging because it can expose some information that can be elided when being printed out, e.g. nil being displayed as nulls, display of quotees, etc. This can be useful in testing.

    fmt.Println(json.MarshalToString(someInterface))
    
The above would print a JSON version of `someInterface{}` as a string.

MarshalToString wraps a call to Marshal. If an error occurs, the error information is discarded an an empty string, `""`, is returned. If you need the error information, call MarshalIndent() instead. Otherwise, the *[]byte* is converted to a *string* and returned.

## Added StringMarshaller struct
The `StringMarshaller` struct provides an easy, compact way of using its `customjson`'s `MarshalToString` and `MarshalIndentToString` functions. `StringMarshaller` offers two Get methods that wrap access to `MarshalToString` and `MarshalIndentToString`, making calls to these functions more compact.

To use, call `customjson.NewStringMarshaller()` and a `StringMarshaller` will be returned with its defaults set. To override the defaults, call `StringMarshaller.Indent(string)` and `StringMarshaller.Prefix(string)`. Each method sets their respective unexported variable. Otherwise the default of `""` and `    ` will be used.

Currently there are two methods on `StringMarshaller`:

**Get**_(interface{}) string_: Takes an interface and returns it as a JSON object converted to a string.
**GetIndented**_(interface{}) string_: Takes an interface and returns it as a JSON object coverted to a formatted string, i.e. indented.

Example:
	package main

	import (
		"fmt"
	
		json "github.com/mohae/customjson"
	)
	
	type team struct {
		City	string
		Name	string
		Mascot	string
	}

	func  main() {

		//Create a new StringMarshaller
		marshalString := json.NewStringMarshaller()

		bulls := &team{City: "Chicago", Name: "Bulls", Mascot: "Benny the Bull"}			
		fmt.Println(marshalString.Get(bulls))
		fmt.Println(marshalString.GetIndented(bulls))
	}

## Notes
I am reconsidering the elision of the HTMLEscape function as it doesn't affect the behavior that this fork of the `json` package was supposed to address. Since this is an exported function that others can use, when needed, it seems useful to leave it in.

Onl considering because I have higher priority code atm. It will be done at some point in the future. This will be signified by both a related commit and revision of this document.

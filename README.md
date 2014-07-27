CustomJSON
==========

A fork of Go's encoding/json package for customization.

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
MarshallIndentToString provides a convenient way of converting something to JSON and then returning a *string* version of it. Since MarshallIndentToString only has a single return value, which is a *string*, it is easy to use whenever you want an formatted string version of any *interface{}*, e.g. printing out a struct value.

    fmt.Println(json.MarshallIndentToString(someInterface, "", "        "))
    
The above would print a formatted version of `someInterface{}`.

MarshallIndentToString wraps a call to MarshallIndent. If an error occurs, the error information is discarded an an empty string, `""`, is returned. If you need the error information, call MarshallIndent() instead. Otherwise, the *[]byte* is converted to a *string* and returned.

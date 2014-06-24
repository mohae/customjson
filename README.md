CustomJSON
==========

Custom JSON object for Go

Main purpose, at time of creation, is to provide a marshaller that doesn't HTMLEscape the marshalled bytes. While this is useful for HTML and browser usage, this is neither helpful or useful in some situations: the marshaller is useful in those situations.

Rather than hack some solution, it seemed better to clone the package and patch it for the desired behavior.

All code contained in this package, with the exception of the modifications necessary to elide the HTML escaping of <,>,&, was originally written by The Go Authors and is licensed using BSD. Please refer to the LICENSE file for more information.

Usage: Same as the Go encode/json package. Use when you don't want the marshalled JSON to be HTMLEscaped. Please note that the JSON output from this package is not safe for use in browsers or use within the HTML \<script\> tags. For those scenarios, Go's enconding/json package should be used as it is designed for those use cases.

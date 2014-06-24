CustomJSON
==========

Custom JSON object for Go

Main purpose, at time of creation, is to provide a marshaller that doesn't HTMLEscape the marshalled bytes. While this is useful for HTML and browser usage, this is neither helpful or useful in some situations: the marshaller is useful in those situations.

Rather than hack some solution, it seemed better to clone the package and elide the code leading to the undesired behavior.

All code contained in this package was originally written by The Go Authors and is licensed using BSD. Please refer to the LICENSE file for more information.

Usage: Same as the Go encode/json package. Use when you don't want the marshalled JSON to be HTMLEscaped.

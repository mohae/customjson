CustomJSON
==========

A fork of Go's encoding/json package.

Main purpose, at time of creation, is to provide a marshaller that doesn't HTMLEscape the marshalled bytes. While this is useful for HTML and browser usage, this is neither helpful or useful in some situations. One example are (Packer)[www.packer.io] templates, which may have inlined bash commands. HTMLEscaping the bash commands is not desirable, as all of the original encoding values must be preserved.

Rather than hack some solution, it seemed better to clone the package and patch it for the desired behavior.

All code contained in this package, with the exception of the modifications necessary to elide the HTML escaping of <,>,&, along with the updating of affected tests, were originally written by The Go Authors.

Usage: Same as the Go encode/json package. Use when you don't want the marshalled JSON to be HTMLEscaped. Please note that the JSON output from this package is not safe for use in browsers or use within the HTML \<script\> tags. For those scenarios, Go's enconding/json package should be used as it is designed for those use cases.

CustomJSON
==========

Custom JSON object for Go

Main purpose, at time of creation, is to provide a marshaller that doesn't HTMLEscape the marshalled bytes. While this is useful for HTML and browser usage, this is neither helpful or useful in some situations: the marshaller is useful in those situations.

# msgp-well-known-types

Well-known types useful for sharing msgp primitives

These currently cover all of go's primitives with the exception of `uintptr`, which msgp doesn't know how to handle.

The intent is not to require that all structs are recursively encodable; it is to ensure that in the simple case, basic primitive messages can be interchanged with minimal overhead.

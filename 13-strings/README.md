# Strings

## Notes

- Strings in Go are different in implementation when compared to other languages.
- A string is a slice of bytes
- Strings in Go are Unicode compliant and are UTF-8 Encoded.
- ```len(s)``` return length of bytes, ```RuneCountInString(s string) (n init)``` can be used to find the length of string.
- ```==``` used to compare 2 strings
- string is immutable. To workaround, string is converted to a slice of runes. Then slice is mutable and converted back to a new string when needed.

## References
- https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/

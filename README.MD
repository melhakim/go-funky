### Go funky
A collection of go functions

JsonToObj
---
A function that recursively converts a json string to a map of objects.
These objects can either be a `string`, `int`, `bool` or `nil`
Useful to convert a json string with an unknown structure.
Depends on `json/encoding`'s `Unmarshal`

```
func JsonToObj(jsonString []byte) map[string]interface{}

```

dencode
========

unified (de|en)coding wrapper for trailing formats.

* toml
* yaml
* json

usage
------

### get dencode

```sh
go get -u github.com/0x75960/dencode
```
### use in code

```go
type TestItem struct {
	A string
	B string
}

data := TestItem{A: "kanasimi", B: "yeeei"}

// toml
dencode.NewEncoder(
	dencode.TomlFormat,
	f,
	).Encode(&data)

dencode.NewDecoder(
	dencode.TomlFormat,
	f,
	).Decode(&data)

// yaml

dencode.NewEncoder(
	dencode.YamlFormat,
	f,
	).Encode(&data)

dencode.NewDecoder(
	dencode.YamlFormat,
	f,
	).Decode(&data)
```
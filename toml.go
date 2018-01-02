package dencode

import (
	"io"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// NewTomlEncoder from io.Writer
func NewTomlEncoder(w io.Writer) (e *toml.Encoder) {
	return toml.NewEncoder(w)
}

// TomlDecoder object
type TomlDecoder struct {
	r io.Reader
}

// NewTomlDecoder from io.Reader
func NewTomlDecoder(r io.Reader) (d *TomlDecoder) {
	return &TomlDecoder{r: r}
}

// Decode a item
func (t TomlDecoder) Decode(out interface{}) (err error) {
	b, err := ioutil.ReadAll(t.r)
	if err != nil {
		return err
	}

	return toml.Unmarshal(b, out)
}

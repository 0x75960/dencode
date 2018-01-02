package dencode

// YamlEncoder
import (
	"bytes"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type YamlEncoder struct {
	w io.Writer
}

// NewYamlEncoder from io.Writer
func NewYamlEncoder(w io.Writer) (e *YamlEncoder) {
	return &YamlEncoder{w: w}
}

// Encode as yaml
func (y *YamlEncoder) Encode(c interface{}) (err error) {

	// yaml.v2 does not have Encoder
	// so implement Encoder roughly.
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	in := bytes.NewBuffer(b)
	_, err = io.Copy(y.w, in)

	return
}

// YamlDecoder object
type YamlDecoder struct {
	r io.Reader
}

// NewYamlDecoder from io.Reader
func NewYamlDecoder(r io.Reader) (d *YamlDecoder) {
	return &YamlDecoder{r: r}
}

func (y *YamlDecoder) Decode(c interface{}) (err error) {

	// yaml.v2 does not have Decoder
	// so implement Decoder roughly.

	b, err := ioutil.ReadAll(y.r)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, c)
}

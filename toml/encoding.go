package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type tomlCodec struct{}

type tomlEncoder struct {
	encoder *toml.Encoder
}

type tomlDecoder struct {
	decoder io.Reader
}

func (c *tomlCodec) NewCodec() encoding.Codec {
	return &tomlCodec{}
}

func (c *tomlCodec) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "codec toml")
}

func (c *tomlCodec) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "codec toml")
}

func (c *tomlCodec) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	return &tomlEncoder{encoder: toml.NewEncoder(w)}, nil
}

func (c *tomlCodec) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(c, v)
}

func (e *tomlEncoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}

func (c *tomlCodec) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &tomlDecoder{decoder: r}, nil
}

func (c *tomlCodec) Decode(b []byte, v interface{}) error {
	return encoding.Decode(c, b, v)
}

func (d *tomlDecoder) Decode(v interface{}) error {
	_, err := toml.DecodeReader(d.decoder, v)
	return err
}

func init() {
	encoding.Register("toml", &tomlCodec{})
}

package importer

var defaultAnnotations = map[string]map[string]string{
	"os": {
		"Stdin":         `*File`,
		"Stdout":        `*File`,
		"Stderr":        `*File`,
		"Create":        `func(name string) (*File \ error)`,
		"(*File).Read":  `func(b []byte) (n int, err error)`,
		"(*File).Write": `func(b []byte) (n int, err error)`,
	},
	"io": {
		"Reader.Read":  `func([]byte) (int, error)`,
		"Writer.Write": `func([]byte) (int, error)`,
	},
	"os/exec": {
		"Command": `func (name string, arg ...string) *Cmd`,
	},
	"html/template": {
		"New":  `func(name string) *Template`,
		"Must": `func(t ?*Template, err ?error) *Template`,
	},
	"strings": {
		"NewReader":      `func(s string) *Reader`,
		"(*Reader).Read": `func(b []byte) (n int, err error)`,
	},
	"go/scanner": {
		"ErrorList": `[]*Error`,
	},
	"errors": {
		"New": `func(text string) error`,
	},
	"net/http": {
		"PostForm":             `func(url string, data url.Values) (resp *Response \ err error)`,
		"HandleFunc":           `func(pattern string, handler func(ResponseWriter, *Request))`,
		"Request.URL":          `*url.URL`,
		"ResponseWriter.Write": `func([]byte) (int, error)`,
		"NewRequest":           `func(method, urlStr string, body ?io.Reader) (*Request \ error)`,
		"(*Client).Do":         `func(req *Request) (resp *Response \ err error)`,
	},
	"encoding/json": {
		"NewDecoder":                `func(io.Reader) *Decoder`,
		"NewEncoder":                `func(io.Writer) *Encoder`,
		"Marshaler.MarshalJSON":     `func() ([]byte \ error)`,
		"Unmarshaler.UnmarshalJSON": `func([]byte) ?error`,
		"Marshal":                   `func(v interface{}) ([]byte \ error)`,
		"Unmarshal":                 `func(data []byte, v ?interface{}) ?error`,
	},
	"flag": {
		"String": `func(name string, value string, usage string) *string`,
		"Usage":  `func()`,
	},
	"fmt": {
		"Errorf": `func(format string, a ...interface{}) error`,
	},
	"bytes": {
		"(*Buffer).Read":  `func(p []byte) (n int, err error)`,
		"(*Buffer).Write": `func(p []byte) (n int, err error)`,
	},
	"time": {
		"Tick": `func(Duration) chan Time`,
	},
}

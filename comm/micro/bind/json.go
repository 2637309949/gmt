package bind

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var JSONSerializer = new(DefaultJSONSerializer)

// DefaultJSONSerializer implements JSON encoding using encoding/json.
type DefaultJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d DefaultJSONSerializer) Serialize(w io.Writer, i interface{}, indent string) error {
	enc := json.NewEncoder(w)
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d DefaultJSONSerializer) Deserialize(r io.Reader, i interface{}) error {
	err := json.NewDecoder(r).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}

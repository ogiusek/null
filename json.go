package null

import (
	"bytes"
	"encoding/json"
)

func (o *Nullable[T]) MarshalJSON() ([]byte, error) {
	if !o.ok {
		return nil, nil
	}
	return json.Marshal(o.val)
}

func (o *Nullable[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nil) || bytes.Equal(data, []byte("null")) {
		*o = Null[T]()
		return nil
	}

	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	*o = New(val)
	return nil
}

package null

import (
	"database/sql/driver"
	"fmt"
)

func (v Nullable[T]) Value() (driver.Value, error) {
	if !v.ok {
		return nil, nil
	}
	return v.val, nil
}

func (v *Nullable[T]) Scan(value interface{}) error {
	if value == nil {
		*v = Null[T]()
		return nil
	}

	scanner, ok := any(v.val).(scanner)
	if ok {
		return scanner.Scan(value)
	}

	val, ok := value.(T)
	if ok {
		*v = New(val)
		return nil
	}

	return fmt.Errorf("cannot convert %v to type %T", value, val)
}

type scanner interface {
	Scan(interface{}) error
}

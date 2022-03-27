package secret

import (
	"encoding/json"
	"reflect"
)

// Make is the standard way that secrets must be created to be properly
// initialized and sealed.
func Make[T any](value T) Secret[T] {
	ret := Secret[T]{}

	// This isn't particulary efficient, however it does seem to prevent the
	// value from ever being printed unintentionally.
	s := &seal[T]{value: value}

	ret.getter = s.get
	ret.setter = s.set

	return ret
}

// Secret contains a value that can be unmarshaled but should never be printed
type Secret[T any] struct {
	getter func() T
	setter func(T)
}

// Get returns the hidden secret value
func (s Secret[T]) Get() T {
	if s.getter == nil {
		panic("using uninitialized secret")
	}

	return s.getter()
}

// Set updates the hidden secret value
func (s Secret[T]) Set(value T) {
	if s.setter == nil {
		panic("using uninitialized secret")
	}

	s.setter(value)
}

// String shows the type of the secret without showing the value.
func (s Secret[T]) String() string {
	return reflect.TypeOf(s).String()
}

// MarshalJSON returns a marshaled form of the underlying secret.  For example,
// a string secret will return an empty string and a numeric secret will return
// a zero value.
func (s Secret[T]) MarshalJSON() ([]byte, error) {
	var ret T
	return json.Marshal(ret)
}

// UnmarshalJSON decodes the input into the underlying secret's value
func (s *Secret[T]) UnmarshalJSON(input []byte) error {
	var dec T
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	*s = Make(dec)

	return nil
}

// seal stores the contents of a secret on the heap and provides accessor
// functions to break reflection-based methods of finding or accidentally print
// the secret.
type seal[T any] struct {
	value T
}

func (s *seal[T]) get() T {
	return s.value
}

func (s *seal[T]) set(value T) {
	s.value = value
}

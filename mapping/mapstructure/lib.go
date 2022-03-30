package mapstructure

import (
	"reflect"

	secret "github.com/andrewbenton/go-secrets"

	"github.com/mitchellh/mapstructure"
)

// DecodeSecretHook handles the translation from a type, T, to a Secret[T]
func DecodeSecretHook[T any]() mapstructure.DecodeHookFunc {
	return mapstructure.DecodeHookFunc(func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		var typ T

		if f != reflect.TypeOf(typ) {
			return data, nil
		}

		if t != reflect.TypeOf(secret.Secret[T]{}) {
			return data, nil
		}

		return secret.Make(data.(T)), nil
	})
}

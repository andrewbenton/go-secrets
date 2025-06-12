package mapstructurev2

import (
	"reflect"

	secret "github.com/andrewbenton/go-secrets"

	mapstructure "github.com/go-viper/mapstructure/v2"
)

func DecodeSecretHook[T any]() mapstructure.DecodeHookFunc {
	return mapstructure.DecodeHookFunc(func(f reflect.Type, t reflect.Type, data any) (any, error) {
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

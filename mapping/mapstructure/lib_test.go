package mapstructure

import (
	"testing"

	secret "github.com/andrewbenton/go-secrets"
	"github.com/stretchr/testify/require"

	"github.com/mitchellh/mapstructure"
)

func TestDecodeSecretHook(t *testing.T) {
	t.Run("string pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": "very-secret",
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.Equal(t, "very-secret", dst.Secret.Get())
	})

	t.Run("string fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": 3,
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})

	t.Run("int pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[int] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[int](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": 3,
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.Equal(t, 3, dst.Secret.Get())
	})

	t.Run("int fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[int] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[int](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": "very-secret",
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})

	t.Run("float pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[float64] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[float64](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": 3.5,
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.Equal(t, 3.5, dst.Secret.Get())
	})

	t.Run("float fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[float64] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[float64](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": "very-secret",
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})

	t.Run("array pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[[]string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[[]string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": []string{
				"a",
				"b",
				"c",
			},
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.EqualValues(t, []string{"a", "b", "c"}, dst.Secret.Get())
	})

	t.Run("array fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[[]string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[[]string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": []int{0, 1, 2},
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})

	t.Run("map pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[map[string]string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[map[string]string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": map[string]string{
				"a": "aay",
				"b": "bee",
				"c": "cee",
			},
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.Equal(t, map[string]string{"a": "aay", "b": "bee", "c": "cee"}, dst.Secret.Get())
	})

	t.Run("map fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[map[string]string] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[map[string]string](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": "not-a-map",
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})

	type tmp struct {
		A int    `mapstructure:"a"`
		B string `mapstructure:"b"`
	}

	t.Run("struct pass", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[tmp] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[tmp](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": tmp{0, "test"},
		}

		err = dec.Decode(input)

		require.NoError(t, err)

		require.Equal(t, tmp{0, "test"}, dst.Secret.Get())
	})

	t.Run("struct fail", func(t *testing.T) {
		var dst struct {
			Secret secret.Secret[tmp] `mapstructure:"secret"`
		}

		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: DecodeSecretHook[tmp](),
			Result:     &dst,
		})

		require.NoError(t, err)

		input := map[string]interface{}{
			"secret": "very-secret",
		}

		err = dec.Decode(input)

		require.Error(t, err)
	})
}

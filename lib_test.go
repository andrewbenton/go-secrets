package secret

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Inner string `json:"inner"`
}

func TestSecret_MarshalJSON(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Parallel()

		s := Make(3)

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "0", string(data))
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()

		s := Make[float64](3)

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "0", string(data))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		s := Make("super secret value")

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "\"\"", string(data))
	})

	t.Run("array", func(t *testing.T) {
		t.Parallel()

		s := Make([]string{"a", "b", "c"})

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "null", string(data))
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()

		s := Make(map[string]string{"a": "aay", "b": "bee", "c": "cee"})

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "null", string(data))
	})

	t.Run("any", func(t *testing.T) {
		t.Parallel()

		s := Make[any]("a")

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, "null", string(data))
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()

		s := Make(testStruct{Inner: "test"})

		data, err := json.Marshal(s)

		require.NoError(t, err)
		require.Equal(t, `{"inner":""}`, string(data))
	})
}

func TestSecret_UnmarshalJSON(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		t.Parallel()

		input := "3"

		var s Secret[int]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.Equal(t, 3, s.Get())
	})

	t.Run("float", func(t *testing.T) {
		t.Parallel()

		input := "3.2"

		var s Secret[float64]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.Equal(t, 3.2, s.Get())
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		input := `"testing"`

		var s Secret[string]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.Equal(t, "testing", s.Get())
	})

	t.Run("array", func(t *testing.T) {
		t.Parallel()

		input := "[1, 2, 3]"

		var s Secret[[]int]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.EqualValues(t, []int{1, 2, 3}, s.Get())
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()

		input := `{"a": "aay", "b": "bee", "c": "cee"}`

		var s Secret[map[string]string]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.EqualValues(t, map[string]string{"a": "aay", "b": "bee", "c": "cee"}, s.Get())
	})

	t.Run("any", func(t *testing.T) {
		t.Parallel()

		input := "3.2"

		var s Secret[any]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.Equal(t, 3.2, s.Get())
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()

		input := `{"inner": "testing"}`

		var s Secret[testStruct]
		err := json.Unmarshal([]byte(input), &s)

		require.NoError(t, err)
		require.Equal(t, s.Get().Inner, "testing")
	})
}

func TestSecret_Extras(t *testing.T) {
	t.Run("%v expansion", func(t *testing.T) {
		t.Parallel()

		sv := Make("inner")

		s := fmt.Sprintf("%#v", sv)

		require.NotContains(t, s, "inner")
	})

	t.Run("string output", func(t *testing.T) {
		t.Parallel()

		sv := Make("inner")

		require.Equal(t, "secret.Secret[string]", sv.String())
	})
}

var testString = "12341234-1234-1234-1234-123412341234"

func BenchmarkJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(testString)
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	input := []byte("\"" + testString + "\"")

	var output string

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(input, &output)
	}
}

func BenchmarkSecret_MarshalJSON(b *testing.B) {
	s := Make(testString)

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(s)
	}
}

func BenchmarkSecret_UnmarshalJSON(b *testing.B) {
	input := []byte("\"" + testString + "\"")

	var output Secret[string]

	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(input, &output)
	}
}

package unicodemisc_test

import (
	"log"
	"testing"

	unicodemisc "github.com/djeeno/go-unicode-misc"
)

// go test -cover -v

var fixture = map[string]byte{
	`\u0000`: 0x00,
	`\u007f`: 0x7f,
	`\u00ff`: 0xff,
}

func TestUnicodeEscapeSequence(t *testing.T) {
	key := `\u0000`

	t.Run("unicodemisc.UnicodeEscapeSequence_"+key, func(t *testing.T) {
		expect := []byte(key)
		actual := unicodemisc.UnicodeEscapeSequence(fixture[key])
		for i := range expect {
			if expect[i] != actual[i] {
				log.Println("expect:", string(expect), "actual", string(actual))
				t.FailNow()
			}
		}
	})
}

func TestUnicodeEscapeSequenceString(t *testing.T) {
	key := `\u00ff`

	t.Run("unicodemisc.UnicodeEscapeSequenceString_"+key, func(t *testing.T) {
		expect := key
		actual := unicodemisc.UnicodeEscapeSequenceString(fixture[key])
		if expect != actual {
			t.FailNow()
		}
	})
}

func TestApppendUnicodeEscapeSequence(t *testing.T) {
	key := `\u007f`

	t.Run("unicodemisc.ApppendUnicodeEscapeSequence_with_nil_"+key, func(t *testing.T) {
		expect := key
		actual := string(unicodemisc.ApppendUnicodeEscapeSequence(nil, fixture[key]))
		if expect != actual {
			t.FailNow()
		}
	})

	t.Run("unicodemisc.ApppendUnicodeEscapeSequence_with_byteSlice_"+key, func(t *testing.T) {
		byteSlice := []byte("head\n")
		expect := string(byteSlice) + key
		actual := string(unicodemisc.ApppendUnicodeEscapeSequence(byteSlice, fixture[key]))
		if expect != actual {
			t.FailNow()
		}
	})
}

// go test -bench . -benchmem -test.run=none -test.benchtime=1000ms

func Benchmark(b *testing.B) {
	for i := 0; i < 5; i++ {
		b.Run("unicodemisc.ApppendUnicodeEscapeSequence", func(b *testing.B) {
			var byteSlice []byte

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				byteSlice = unicodemisc.ApppendUnicodeEscapeSequence(byteSlice, 0x00)
			}
			_ = byteSlice
		})

		b.Run("unicodemisc.UnicodeEscapeSequence", func(b *testing.B) {
			var byteSlice []byte

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				byteSlice = append(byteSlice, unicodemisc.UnicodeEscapeSequence(0x00)...)
			}
			_ = byteSlice
		})
	}
}

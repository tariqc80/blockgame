package block

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/tariqc80/blockgame/pkg/util"
)

func TestData(t *testing.T) {
	data := NewData("input text", "HASHVALUE")

	t.Run("`NewData`", func(t *testing.T) {
		if data == nil {
			t.Fatal("`NewData` return nil value")
		}
	})

	t.Run("`sum`", func(t *testing.T) {
		raw := make([]byte, 0)

		raw = append(raw, util.Uint64ToHexBytes(data.id)...)
		raw = append(raw, []byte(data.text)...)
		raw = append(raw, []byte(data.previousHash)...)
		raw = append(raw, util.Int64ToHexBytes(data.timestamp)...)

		hashed := sha256.Sum256(raw)
		expected := hex.EncodeToString(hashed[:])

		got := data.sum()

		if got != expected {
			t.Errorf("Expected hash to be %s, but got %s", expected, got)
		}
	})

	t.Run("`Timestamp`", func(t *testing.T) {
		got := data.Timestamp()

		if got != data.timestamp {
			t.Errorf("Expected `Timestamp` to return %d", data.timestamp)
		}
	})
}

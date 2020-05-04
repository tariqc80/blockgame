package block

import "testing"

func TestBlock(t *testing.T) {
	blk := NewBlock("Genesis Block", nil)
	next := NewBlock("Second Block", blk)

	t.Run("`NewBlock`", func(t *testing.T) {
		t.Run("with nil previous", func(t *testing.T) {
			if blk == nil {
				t.Fatal("Expected `NewBlock` to return a non-nil value")
			}

			if blk.previous != nil {
				t.Error("Expected `previous` to be nil")
			}

			data, ok := blk.data.(*Data)
			if !ok {
				t.Fatalf("Expected `data` property to be type `*Data` got %T", blk.data)
			}

			if data.text != "Genesis Block" {
				t.Errorf("Expected 'Genesis Block' got %s", data.text)
			}

			if data.previousHash != "" {
				t.Errorf("Expected `previousHash` to be an empty string")
			}
		})

		t.Run("with real previous", func(t *testing.T) {
			if next.previous == nil {
				t.Error("Expected `previous` to be non-nil")
			}

			if next.data.(*Data).previousHash != blk.GetHash() {
				t.Error("Expected `previousHash` to match previous blocks hash")
			}
		})
	})

	t.Run("`Data`", func(t *testing.T) {
		got, ok := blk.Data().(*Data)

		if !ok {
			t.Fatal("Expected `Data` to return `*Data`")
		}

		if got.text != "Genesis Block" {
			t.Error("Unexpected data text value")
		}
	})

	t.Run("`GetHash`", func(t *testing.T) {
		if blk.GetHash() != blk.data.sum() {
			t.Error("Hash does not have checksum for block data")
		}
	})

	t.Run("`Previous`", func(t *testing.T) {
		if next.Previous() != blk {
			t.Error("Unexpected value returned from `Previous()`. Should be pointer to previous block")
		}
	})

	t.Run("`Validate`", func(t *testing.T) {

	})
}

package main

import (
	"fmt"

	"github.com/tariqc80/blockgame/internal/block"
)

func main() {
	fmt.Println("vim-go")

	blk := block.NewBlock("my first block", nil)

	fmt.Println(blk.Data())

	fmt.Println(blk.GetHash())
	fmt.Println(blk.GetHash())

	blk2 := block.NewBlock("second block", blk)

	fmt.Println(blk2.Data())
	fmt.Println(blk2.GetHash())
}

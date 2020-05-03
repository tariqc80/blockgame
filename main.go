package main

import (
	"fmt"

	"github.com/tariqc80/blockgame/internal/block"
)

func main() {
	fmt.Println("vim-go")

	blk := block.NewBlock()

	fmt.Println(blk.Data())
}

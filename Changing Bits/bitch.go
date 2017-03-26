package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var n, q, idx int
	var x uint
	var cmd string
	a, b, c := new(big.Int), new(big.Int), new(big.Int)
	var output bytes.Buffer
	var recalc = true
	input := bufio.NewReader(os.Stdin)
	fmt.Fscan(input, &n)
	fmt.Fscan(input, &q)
	input.ReadRune()
	fmt.Fscanf(input, "%b", a)
	input.ReadRune()
	fmt.Fscanf(input, "%b", b)
	input.ReadRune()
	for i := 0; i < q; i++ {
		fmt.Fscan(input, &cmd)
		fmt.Fscan(input, &idx)
		switch cmd[4] {
		case 'a':
			fmt.Fscan(input, &x)
			a.SetBit(a, idx, x)
			recalc = true
		case 'b':
			fmt.Fscan(input, &x)
			b.SetBit(b, idx, x)
			recalc = true
		case 'c':
			if recalc {
				c.Add(a, b)
				recalc = false
			}
			output.WriteString(fmt.Sprintf("%d", c.Bit(idx)))
		}
	}
	output.WriteTo(os.Stdout)
}

package lib

import (
	"fmt"
	"os"
)

func Arg() {
	argsRaw := os.Args
	fmt.Printf("=>\t %#v\n", argsRaw)

	args := argsRaw[1:]
	fmt.Printf("=>\t %#v\n", args)
}

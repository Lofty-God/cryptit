package main

import (
	"fmt"

	"github.com/Lofty-God/cryptit/decrypt"
	"github.com/Lofty-God/cryptit/encrypt"
)

func main() {
	encryptedstr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedstr)
	fmt.Println(decrypt.Nimbus(encryptedstr))
}

package main

import (
	"fmt"

	"github.com/Lofty-God/cryptit/decrypt"
	"github.com/Lofty-God/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}

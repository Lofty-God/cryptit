package main

import (
	"fmt"

	"github.com/Lofty-God/cryptit/encrypt"

	"github.com/pborman/uuid"
)

func main() {
	encryptedstr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedstr)
	uuid := uuid.NewRandom()
	fmt.Println(uuid)
}

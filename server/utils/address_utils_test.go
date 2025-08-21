package utils

import (
	"fmt"
	"testing"
)

func TestIsValidCryptoAddress(t *testing.T) {

	eth_address := "0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3 "

	tron_address := "TS4WHd3PyEiYXDxRZbmofj1zugudW6Dior"
	flag1 := IsValidTronAddress(tron_address)

	flag2 := IsValidEthereumAddress(eth_address)

	fmt.Println(flag1, flag2)

	result1, message1 := IsValidCryptoAddress(eth_address)
	result2, message2 := IsValidCryptoAddress(tron_address)

	fmt.Println(result1, message1)

	fmt.Println(result2, message2)
}

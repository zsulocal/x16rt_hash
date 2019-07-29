package x16rt_hash

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestX16r(t *testing.T) {
	hash := "000000209f6a3da85c6f63ead1df53a923410838929fbaa3ffbf9dfa74f4b46af7cff872f5ec36b7df9b25eca7305456d38337a66c09a19960fe83caeb9cef98a5de0b9600ea3e5d6eb50e1b082f2a4a0000"
	hash_bin, _ := hex.DecodeString(hash)
	powhash := GetPowHash(hash_bin)
	powhash_hex := hex.EncodeToString(powhash)
	fmt.Println(powhash_hex)
	if powhash_hex != "786d2422983e1b7b3baad5bd076967088b4966077d83fae99c04060000000000" {
		t.Error("Test x16r powhash failed")
		return
	}
}

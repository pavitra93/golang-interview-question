package main

import (
	"fmt"
	"strings"
)

func main() {
	var Codec Codec
	encode := Codec.Encode([]string{"hello", "world"})
	fmt.Println(encode)
	fmt.Println(Codec.Decode(encode))
}

type Codec struct {
}

// Encode Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var EncodeString string
	EncodeString = strings.Join(strs, "π")

	return EncodeString

}

// Decode Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var DecodeString []string
	DecodeString = strings.Split(strs, "π")

	return DecodeString
}

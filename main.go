package main

import (
	"fmt"

	"github.com/rylans/getlang"
)

func main() {
	info := getlang.FromString("How are you now?")
	fmt.Println(info.LanguageCode(), info.LanguageName(), info.Confidence())
}

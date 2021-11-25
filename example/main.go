package main

import (
	"fmt"
	"github.com/go-curses/bitmasker/example/mental"
)

func main() {
	fmt.Println("# State")
	s := mental.Unconscious
	fmt.Printf("0: %s\n", s)
	s = s.Set(mental.Conscious)
	fmt.Printf("1: %s\n", s)
	s = s.Set(mental.Meditative)
	fmt.Printf("2: %s\n", s)
	s = s.Set(mental.Entertained).Clear(mental.Meditative)
	fmt.Printf("3: %s\n", s)
	fmt.Println("# ThingFlags")
	t := mental.ThingZeroFlag
	fmt.Printf("0: %s\n", t)
	t = t.Set(mental.ThingOneFlag)
	fmt.Printf("1: %s\n", t)
	t = t.Set(mental.ThingTwoFlag)
	fmt.Printf("2: %s\n", t)
	t = t.Clear(mental.ThingOneFlag)
	fmt.Printf("3: %s\n", t)
}

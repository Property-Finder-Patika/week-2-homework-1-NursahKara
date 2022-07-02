package main

func main() {

	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	uint8Val = byteVal

	var (
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val

	runeVal = rune(int32Val)
	runeVal = rune(runeVal)

	_, _, _, _ = byteVal, uint8Val, intVal, runeVal
}

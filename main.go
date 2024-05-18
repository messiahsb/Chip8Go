package main

import (
	//"bytes"
	"fmt"
	// "internal/cpu"
	"log"
	"os"
	// "time"
)

type Chip8 struct {
	opcode      uint16
	memory      [4096]byte
	V           [16]byte
	I           uint16
	PC          uint16
	graphics    [64 * 32]byte
	delay_timer byte
	sound_timer byte
	stack       [16]uint16
	sp          uint16
	key         [16]byte
}

func (chip *Chip8) handleOpcode() {

	// chip.opcode = uint16(chip.memory[chip.PC] << 4)
	firstNib := uint16(chip.memory[chip.PC])

	fmt.Printf("\n This is the first nib: %X\n", firstNib)

	switch firstNib {
	case 0x00:
		fmt.Println("0 not handled yet")
	case 0x01:
		fmt.Println("1 not handled yet")
	case 0x02:
		fmt.Println("2 not handled yet")
	case 0x03:
		fmt.Println("3 not handled yet")
	case 0x04:
		fmt.Println("4 not handled yet")
	case 0x05:
		fmt.Println("5 not handled yet")
	case 0x06:
		fmt.Println("6 not handled yet")
	case 0x07:
		fmt.Println("7 not handled yet")
	case 0x08:
		fmt.Println("8 not handled yet")
	default:
		fmt.Println("Opcode unknown")
	}

	// if chip.delay_timer > 0 {
	// 	chip.delay_timer--
	// }
	// if chip.sound_timer > 0 {
	// 	if chip.sound_timer == 1 {
	// 		fmt.Println("BEEP!")
	// 	}
	// 	chip.sound_timer--
	// }
}

func main() {
	//setup graphics
	//setup input

	chip := Chip8{}
	file, err := os.Open("ibm_logo.ch8") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	// data := make([]byte, 1000)
	count, err := file.Read(chip.memory[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, chip.memory[:count])

	chip.handleOpcode()

	//insitialize chip8chip and load rom
	//c8 := Chip8{}

	//for loop for emulation

	//emulate one cycle

	//update screen

	//store keypress

}

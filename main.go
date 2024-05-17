package main

import (
	//"bytes"
	"fmt"
	// "internal/cpu"
	// "log"
	// "os"
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

func (cpu *Chip8) emulateCycle() {

	cpu.opcode = uint16(cpu.memory[cpu.PC] << 4)

	switch cpu.opcode {
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

		if cpu.delay_timer > 0 {
			cpu.delay_timer--
		}
		if cpu.sound_timer > 0 {
			if cpu.sound_timer == 1 {
				fmt.Println("BEEP!")
			}
			cpu.sound_timer--
		}
	}
}

func main() {
	//setup graphics
	//setup input

	//insitialize chip8cpu and load rom
	//c8 := Chip8{}

	//for loop for emulation

	//emulate one cycle

	//update screen

	//store keypress

}

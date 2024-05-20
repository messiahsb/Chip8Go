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

	chip.opcode = uint16(chip.memory[chip.PC])<<8 | uint16(chip.memory[chip.PC+1])
	// firstNib := uint8(chip.memory[chip.PC] >> 4)
	firstNib := chip.opcode & 0xf000
	fmt.Printf("\n This is the opcode: %X\n", (chip.opcode))
	fmt.Printf("\n This is the first nib: %X\n", firstNib)
	//fmt.Printf("\n This is the second nib: %X\n", (chip.memory[chip.PC] & 0x0F))

	switch firstNib {
	// 1110
	//000N Call, 00E0 Dislplay clear, 00EE return
	case 0x0000:
		fmt.Println("0 not handled yet")
	//1NNN goto address NNN
	case 0x1000:
		fmt.Println("1 not handled yet")
	//2NNN calls subroutine at NNN
	case 0x2000:
		fmt.Println("2 not handled yet")
	//3XNN if Vx == NN
	case 0x3000:
		fmt.Println("3 not handled yet")
	//4XNN if Vx != NN
	case 0x4000:
		fmt.Println("4 not handled yet")
	//5XY0 if Vx == Vy
	case 0x5000:
		fmt.Println("5 not handled yet")
	//6XNN set vx = nn
	case 0x6000:
		VX := (chip.opcode & 0x0F00) >> 8
		chip.V[VX] = uint8(chip.opcode & 0x00FF)
		fmt.Printf("6: tentative, print vx: %X\n", chip.V[VX])
	//7XNN adds NN to Vx
	case 0x7000:
		VX := (chip.opcode & 0x0F00) >> 8
		chip.V[VX] += uint8(chip.opcode & 0x00FF)
		fmt.Printf("7: tentative, print vx: %X\n", chip.V[VX])
	//8XY
	case 0x8000:
		//add switch case for different 8xxx values
		VX := (chip.opcode & 0x0F00) >> 8
		VY := (chip.opcode & 0x00F0) >> 8
		chip.V[VX] = chip.V[VY]
		fmt.Printf("8: tentative, print vx: %X\n and print vy: %X\n", chip.V[VX], chip.V[VY])

	case 0x9000:
		fmt.Println("0 not handled yet")
	case 0xA000:
		chip.I = chip.opcode & 0x0FFF
		fmt.Printf("A: tentative, print I: %X\n", chip.I)

	case 0xB000:
		fmt.Println("B not handled yet")
	case 0xC000:
		fmt.Println("C not handled yet")
	case 0xD000:
		fmt.Println("D not handled yet")
	case 0xE000:
		fmt.Println("E not handled yet")
	case 0xF000:
		fmt.Println("F not handled yet")
	default:
		fmt.Println("Opcode unknown")
	}
	chip.PC += 1
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

	//intializes cpu
	chip := Chip8{}
	//opens rom input
	file, err := os.Open("ibm_logo.ch8") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	// fileinfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Size of the file:", fileinfo.Size())

	// data := make([]byte, 1000)

	//loads rom to cpu mem
	count, err := file.Read(chip.memory[:])
	if err != nil {
		log.Fatal(err)
	}
	//  fmt.Printf("read %d bytes: %q\n", count, chip.memory[:count])
	fmt.Printf("read %d ", count)

	//intialize for loop to run program
	// chip.PC = 200
	for chip.PC < 10 {

		//emulate one cycle
		chip.handleOpcode()
		chip.PC = chip.PC + 1
		//update screen

		//store keypress
	}
}

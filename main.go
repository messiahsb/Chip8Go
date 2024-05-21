package main

import (
	//"bytes"
	"fmt"
	// "internal/cpu"
	"math/rand/v2"
	"os"

	// "time"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

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
		chip.PC += 2
		fmt.Println("0 not handled yet")
		//switc
	//1NNN goto address NNN
	case 0x1000:
		chip.PC = chip.opcode & 0x0FFF
		fmt.Println("1	handled")
	//2NNN calls subroutine at NNN
	case 0x2000:
		chip.PC += 2
		fmt.Println("2 not handled yet")
	//3XNN if Vx == NN
	case 0x3000:
		VX := (chip.opcode & 0x0F00) >> 8
		if chip.V[VX] == byte(chip.opcode&0x00FF) {
			chip.PC += 4
		} else {
			chip.PC += 2
		}
		fmt.Println("3 handled")
	//4XNN if Vx != NN
	case 0x4000:
		VX := (chip.opcode & 0x0F00) >> 8
		if chip.V[VX] != byte(chip.opcode&0x00FF) {
			chip.PC += 4
		} else {
			chip.PC += 2
		}
		fmt.Println("4 handled")
	//5XY0 if Vx == Vy
	case 0x5000:
		VX := (chip.opcode & 0x0F00) >> 8
		VY := (chip.opcode & 0x00F0) >> 8
		if chip.V[VX] == chip.V[VY] {
			chip.PC += 4
		} else {
			chip.PC += 2
		}
		fmt.Println("5 Handled")
	//6XNN set vx = nn
	case 0x6000:
		//VX gets the target register by bitwise operation for the second nibble
		VX := (chip.opcode & 0x0F00) >> 8
		//gets the value the register is being set to from the last two bytes and sets it
		chip.V[VX] = uint8(chip.opcode & 0x00FF)
		chip.PC += 2
		fmt.Printf("6: tentative, print vx: %X\n", chip.V[VX])
	//7XNN adds NN to Vx
	case 0x7000:
		VX := (chip.opcode & 0x0F00) >> 8
		chip.V[VX] += uint8(chip.opcode & 0x00FF)
		chip.PC += 2
		fmt.Printf("7: tentative, print vx: %X\n", chip.V[VX])
	//8XY
	case 0x8000:
		//added switch case for different 8xxx values
		switch (chip.opcode & 0x000F) >> 8 {
		case 0x0:
			//sets VX to value of VY
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			chip.V[VX] = chip.V[VY]
			chip.PC += 2
		case 0x1:
			//bitwise vx or vy
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			chip.V[VX] |= chip.V[VY]
			chip.PC += 2
		case 0x2:
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			//bitwise vx and vy
			chip.V[VX] &= chip.V[VY]
			chip.PC += 2
		case 0x3:
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			// vx xor vy
			chip.V[VX] ^= chip.V[VY]
			chip.PC += 2
		case 0x4:
			// add vy to vx, vf set to 1 if theres overflow and 0 if not
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			if chip.V[VY] > (0xFF - chip.V[VX]) {
				chip.V[0xF] = 1
			} else {
				chip.V[0xF] = 0
			}
			chip.V[VX] += chip.V[VY]
			chip.PC += 2
		case 0x5:
			// subtract vy from vx, vf set to 1 if theres underflow and 0 if not
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			if chip.V[VX] >= chip.V[VY] {
				chip.V[0xF] = 1
			} else {
				chip.V[0xF] = 0
			}
			chip.V[VX] -= chip.V[VY]
			chip.PC += 2
		case 0x6:
			// Stores the least significant bit of VX in VF and then shifts VX to the right by 1
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			chip.V[0xF] = chip.V[VY] & 0x01
			chip.V[VX] = chip.V[VY] >> 1
			chip.PC += 2
		case 0x7:
			//Sets VX to VY minus VX. VF is set to 0 when there's an underflow, and 1 when there is not. (i.e. VF set to 1 if VY >= VX)
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			if chip.V[VY] >= chip.V[VX] {
				chip.V[0xF] = 1
			} else {
				chip.V[0xF] = 0
			}
			chip.V[VX] = chip.V[VY] - chip.V[VX]
			chip.PC += 2
		case 0xE:
			//Stores the most significant bit of VX in VF and then shifts VX to the left by 1
			VX := (chip.opcode & 0x0F00) >> 8
			VY := (chip.opcode & 0x00F0) >> 8
			chip.V[0xF] = chip.V[VY] & 0x80
			chip.V[VX] = chip.V[VY] << 1
			chip.PC += 2
		default:
			fmt.Println("opknown at 8")
		}
		VX := (chip.opcode & 0x0F00) >> 8
		VY := (chip.opcode & 0x00F0) >> 8
		fmt.Printf("8: tentative, print vx: %X\n and print vy: %X\n", chip.V[VX], chip.V[VY])
	case 0x9000:
		VX := (chip.opcode & 0x0F00) >> 8
		VY := (chip.opcode & 0x00F0) >> 8
		if chip.V[VY] != chip.V[VX] {
			chip.PC += 4
		}
		chip.PC += 2
		fmt.Println("9 handled ")
	case 0xA000:
		chip.I = chip.opcode & 0x0FFF
		chip.PC += 2
		fmt.Printf("A: tentative, print I: %X\n", chip.I)

	case 0xB000:
		chip.PC = (chip.opcode & 0x0FFF) + uint16(chip.V[0x0])
		chip.PC += 2
		fmt.Println("B handled")
	case 0xC000:
		//Sets VX to the result of a bitwise and operation on a random number
		VX := chip.opcode & 0x0FFF
		chip.V[VX] = byte(rand.UintN(255)) & byte(chip.opcode&0x00FF)
		chip.PC += 2
		fmt.Println("C handled")
	case 0xD000:
		//Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
		chip.PC += 2
		fmt.Println("D not handled yet")
	case 0xE000:
		//Skips the next instruction if the key stored in VX is pressed (usually the next instruction is a jump to skip a code block).[22]
		//Skips the next instruction if the key stored in VX is not pressed (usually the next instruction is a jump to skip a code block)
		chip.PC += 2
		fmt.Println("E not handled yet")
	case 0xF000:
		//FXXX to set timer or interact with memory
		chip.PC += 2
		fmt.Println("F not handled yet")
	default:
		fmt.Println("Opcode unknown")
	}

}

func main() {
	//setup graphics
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
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
		// chip.PC = chip.PC + 1
		//update screen

		//store keypress

		if chip.delay_timer > 0 {
			chip.delay_timer--
		}

		if chip.sound_timer > 0 {
			if chip.sound_timer == 1 {
				fmt.Println("BEEP!")
			}
			chip.sound_timer--
		}
	}
}

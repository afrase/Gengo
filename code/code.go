package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Instructions for the code package.
type Instructions []byte

// Opcode is an opcode in the VM.
type Opcode byte

const (
	// OpConstant stores the index of a constant in the constant pool.
	OpConstant Opcode = iota
	// OpPop tells the VM to pop the topmost element off of the stack.
	OpPop
	// OpAdd tells the VM to pop the two topmost elements off the stack,
	// add them together and push the result back on to the stack.
	OpAdd
	// OpSub is the same as OpAdd except for subtraction.
	OpSub
	// OpMul is the same but for multiplication.
	OpMul
	// OpDiv is for division.
	OpDiv
	// OpTrue tells the VM to push a true value onto the stack.
	OpTrue
	// OpFalse tells the VM to push a false value onto the stack.
	OpFalse
	// OpEqual tells the VM to equally compares two topmost elements on the stack.
	OpEqual
	// OpNotEqual tells the VM to perform a not equal comparison between two values.
	OpNotEqual
	// OpGreaterThan tells the VM to perform a greater than comparison between two values.
	OpGreaterThan
	// OpMinus tells the VM to negate an integer.
	OpMinus
	// OpBang tells the VM to negate booleans.
	OpBang
	// OpJumpNotTruthy jumps to the instruction if the value at the stop of the stack is not truthy.
	OpJumpNotTruthy
	// OpJump represents the opcode for jumping to a specific position in the code.
	OpJump
	// OpNull stores the null value in the operand stack.
	OpNull
)

// Definition of each opcode.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:      {"OpConstant", []int{2}},
	OpPop:           {"OpPop", []int{}},
	OpAdd:           {"OpAdd", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpMul", []int{}},
	OpDiv:           {"OpDiv", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEqual:         {"OpEqual", []int{}},
	OpNotEqual:      {"OpNotEqual", []int{}},
	OpGreaterThan:   {"OpGreaterThan", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},
	OpNull:          {"OpNull", []int{}},
}

// Lookup returns the definition for a given opcode.
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make a bytecode
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			_, _ = fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])
		_, _ = fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

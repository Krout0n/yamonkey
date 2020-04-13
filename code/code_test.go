package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)
		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d", len(tt.expected), len(instruction))
		}
		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d", i, b, instruction[i])
			}
		}
	}

}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpConstant, 1),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}
	expected := `
0000 OpConstant 1
0003 OpConstant 2
0006 OpConstant 6553
`[1:]
	concatted := Instructions{}

	for _, ins := range instructions {
		// concatted = [0 0 1 0 0 2 0 255 255]
		// 0 0 1 <- OpConstant 0x00 0x01
		// 0 0 2 <- OpConstant 0x00 0x02
		// 0 255 255 <- OpConstant 0xff 0xff
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
		{OpConstant, []int{1}, 2},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)
		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}
		operandsRead, n := ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d\n", tt.bytesRead, n)
		}

		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operand wrong. want=%d, got=%d", want, operandsRead[i])
			}

		}
	}
}

func TestReadUint16(t *testing.T) {
	tests := []struct {
		arg      []byte
		expected uint16
	}{
		{[]byte{1, 2}, 258},
		{[]byte{1, 2, 3}, 258},
	}

	for _, tt := range tests {
		if got := ReadUint16(Instructions(tt.arg)); tt.expected != got {
			t.Fatalf("expected: %d, got: %d", tt.expected, got)
		}
	}
}

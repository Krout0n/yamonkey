package vm

import (
	"fmt"
	"monkey/code"
	"monkey/compiler"
	"monkey/object"
)

const stackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object
	sp           int
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,
		stack:        make([]object.Object, stackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	// ip is an instruction pointer.
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])
		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
			ip += 2
		case code.OpAdd:
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(object.Integer).Value
			rightValue := right.(object.Integer).Value
			result := leftValue + rightValue
			vm.push(object.Integer{Value: result})
		case code.OpPop:
			vm.pop()
		}

	}
	return nil
}

func (vm *VM) push(obj object.Object) error {
	if stackSize <= vm.sp+1 {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = obj
	vm.sp++
	return nil
}

func (vm *VM) pop() object.Object {
	if vm.sp-1 < 0 {
		return nil
	}
	obj := vm.constants[vm.sp-1]
	vm.sp--
	return obj
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

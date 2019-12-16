package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ISHL struct{ base.NoOperandsInstruction } // int左位移

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f

	result := v1 << s
	stack.PushInt(result)
}

type ISHR struct{ base.NoOperandsInstruction } // int算术右位移

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

/*
int x = -1;
println(Integer.toBinaryString(x)); // 11111111111111111111111111111111
println(Integer.toBinaryString(x >> 8)); // 11111111111111111111111111111111
println(Integer.toBinaryString(x >>> 8)); // 00000000111111111111111111111111
*/

type IUSHR struct{ base.NoOperandsInstruction } // int逻辑右位移

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

type LSHL struct{ base.NoOperandsInstruction } // long左位移

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f

	result := v1 << s
	stack.PushLong(result)
}

type LSHR struct{ base.NoOperandsInstruction } // long算术右位移

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f

	result := v1 >> s
	stack.PushLong(result)
}

type LUSHR struct{ base.NoOperandsInstruction } // long逻辑右位移

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}

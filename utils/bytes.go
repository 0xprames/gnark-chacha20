package utils

import (
	"encoding/binary"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"
)

func BytesToUint32LE(in []uint8) []uints.U32 {

	var res []uints.U32
	for i := 0; i < len(in); i += 4 {
		t := binary.LittleEndian.Uint32(in[i:])
		res = append(res, uints.NewU32(t))
	}
	return res
}
func BytesToUint32BE(in []uint8) []uints.U32 {

	var res []uints.U32
	for i := 0; i < len(in); i += 4 {
		t := binary.BigEndian.Uint32(in[i:])
		res = append(res, uints.NewU32(t))
	}
	return res
}

func BytesToUint32BERaw(in []uint8) []frontend.Variable {

	var res []frontend.Variable
	for i := 0; i < len(in); i += 4 {
		t := binary.BigEndian.Uint32(in[i:])
		res = append(res, t)
	}
	return res
}

func BytesToUint32LERaw(in []uint8) []frontend.Variable {

	var res []frontend.Variable
	for i := 0; i < len(in); i += 4 {
		t := binary.LittleEndian.Uint32(in[i:])
		res = append(res, t)
	}
	return res
}

func Uint32ToBits(in frontend.Variable) [32]frontend.Variable {
	var b *big.Int
	switch in.(type) {
	case uint32:
		b = big.NewInt(int64(in.(uint32)))
		break
	case int:
		b = big.NewInt(int64(in.(int)))
		break
	default:
		panic("invalid type")
	}

	var res [32]frontend.Variable
	for i := 0; i < 32; i++ {
		res[i] = b.Bit(i)
	}
	return res
}

func UintsToBits(in []frontend.Variable) [][32]frontend.Variable {
	res := make([][32]frontend.Variable, len(in))
	for i := 0; i < len(in); i++ {
		res[i] = Uint32ToBits(in[i])
	}
	return res
}

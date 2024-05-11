package encode

import "unsafe"

type Uint64Encoder uint64

func (e Uint64Encoder) Encode() ([]byte, error) {
	return (*(*[8]byte)(unsafe.Pointer(&e)))[:], nil
}

func (e Uint64Encoder) Length() int {
	return 8
}

type Int64Encoder int64

func (e Int64Encoder) Encode() ([]byte, error) {
	return Uint64Encoder(e).Encode()
}

func (e Int64Encoder) Length() int {
	return 8
}

type IntEncoder int

func (e IntEncoder) Encode() ([]byte, error) {
	return Uint64Encoder(e).Encode()
}

func (e IntEncoder) Length() int {
	return 8
}

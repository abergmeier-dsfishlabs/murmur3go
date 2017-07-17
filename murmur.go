
package murmur

import (
	"encoding/binary"
)

func ROTL32(x uint32, r uint8) uint32 {
	return (x << r) | (x >> (32 - r))
}

func fmix32 (h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16

	return h
}

func getblock32 (p []byte, tailIndex int, i int) (result uint32) {
	result = binary.LittleEndian.Uint32(p[tailIndex + (i * 4):])
	return
}

func MurmurHash3_x86_128 (key []byte, seed uint32) (out [2]uint64) {
	nblocks := len(key) / 16
	// tail points just behind the last complete 32 bit block in the key
	tailIndex := 16 * nblocks
	tail := key[tailIndex:]

	h1 := seed
	h2 := seed
	h3 := seed
	h4 := seed

	//----------
	// body

	for i := -nblocks; i != 0; i++ {
		k1 := getblock32(key, tailIndex, i*4+0)
		k2 := getblock32(key, tailIndex, i*4+1)
		k3 := getblock32(key, tailIndex, i*4+2)
		k4 := getblock32(key, tailIndex, i*4+3)

		k1 *= c1; k1  = ROTL32(k1,15); k1 *= c2; h1 ^= k1;

		h1 = ROTL32(h1,19); h1 += h2; h1 = h1*5+0x561ccd1b;

		k2 *= c2; k2  = ROTL32(k2,16); k2 *= c3; h2 ^= k2;

		h2 = ROTL32(h2,17); h2 += h3; h2 = h2*5+0x0bcaa747;

		k3 *= c3; k3  = ROTL32(k3,17); k3 *= c4; h3 ^= k3;

		h3 = ROTL32(h3,15); h3 += h4; h3 = h3*5+0x96cd1c35;

		k4 *= c4; k4  = ROTL32(k4,18); k4 *= c1; h4 ^= k4;

		h4 = ROTL32(h4,13); h4 += h1; h4 = h4*5+0x32ac3b17;
	}

	//----------
	// tail

	k1 := uint32(0)
	k2 := uint32(0)
	k3 := uint32(0)
	k4 := uint32(0)

	switch len(key) & 15 {
	case 15: k4 ^= uint32(tail[14]) << 16; fallthrough
	case 14: k4 ^= uint32(tail[13]) << 8; fallthrough
	case 13: k4 ^= uint32(tail[12]) << 0
		k4 *= c4; k4  = ROTL32(k4,18); k4 *= c1; h4 ^= k4;
		fallthrough
	case 12: k3 ^= uint32(tail[11]) << 24; fallthrough
	case 11: k3 ^= uint32(tail[10]) << 16; fallthrough
	case 10: k3 ^= uint32(tail[ 9]) << 8; fallthrough
	case  9: k3 ^= uint32(tail[ 8]) << 0
		k3 *= c3; k3  = ROTL32(k3,17); k3 *= c4; h3 ^= k3;
		fallthrough
	case  8: k2 ^= uint32(tail[ 7]) << 24; fallthrough
	case  7: k2 ^= uint32(tail[ 6]) << 16; fallthrough
	case  6: k2 ^= uint32(tail[ 5]) << 8; fallthrough
	case  5: k2 ^= uint32(tail[ 4]) << 0
		k2 *= c2; k2  = ROTL32(k2,16); k2 *= c3; h2 ^= k2;
		fallthrough
	case  4: k1 ^= uint32(tail[ 3]) << 24; fallthrough
	case  3: k1 ^= uint32(tail[ 2]) << 16; fallthrough
	case  2: k1 ^= uint32(tail[ 1]) << 8; fallthrough
	case  1: k1 ^= uint32(tail[ 0]) << 0
		k1 *= c1; k1  = ROTL32(k1,15); k1 *= c2; h1 ^= k1;
	}

	//----------
	// finalization
	keyLen := uint32(len(key))
	h1 ^= keyLen; h2 ^= keyLen; h3 ^= keyLen; h4 ^= keyLen;

	h1 += h2; h1 += h3; h1 += h4;
	h2 += h1; h3 += h1; h4 += h1;

	h1 = fmix32(h1)
	h2 = fmix32(h2)
	h3 = fmix32(h3)
	h4 = fmix32(h4)

	h1 += h2; h1 += h3; h1 += h4;
	h2 += h1; h3 += h1; h4 += h1;

	buffer := make([]byte, 8)
	binary.LittleEndian.PutUint32(buffer[:4], h1)
	binary.LittleEndian.PutUint32(buffer[4:], h2)
	out[0] = binary.LittleEndian.Uint64(buffer)
	binary.LittleEndian.PutUint32(buffer[:4], h3)
	binary.LittleEndian.PutUint32(buffer[4:], h4)
	out[1] = binary.LittleEndian.Uint64(buffer)

	return
}


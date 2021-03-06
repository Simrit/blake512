// Written in 2011-2012 by Dmitry Chestnykh.
//
// To the extent possible under law, the author have dedicated all copyright
// and related and neighboring rights to this software to the public domain
// worldwide. This software is distributed without any warranty.
// http://creativecommons.org/publicdomain/zero/1.0/

// BLAKE-512 block step.
// In its own file so that a faster assembly or C version
// can be substituted easily.

package blake512

const (
	cst0  = 0x243F6A8885A308D3
	cst1  = 0x13198A2E03707344
	cst2  = 0xA4093822299F31D0
	cst3  = 0x082EFA98EC4E6C89
	cst4  = 0x452821E638D01377
	cst5  = 0xBE5466CF34E90C6C
	cst6  = 0xC0AC29B7C97C50DD
	cst7  = 0x3F84D5B5B5470917
	cst8  = 0x9216D5D98979FB1B
	cst9  = 0xD1310BA698DFB5AC
	cst10 = 0x2FFD72DBD01ADFB7
	cst11 = 0xB8E1AFED6A267E96
	cst12 = 0xBA7C9045F12C7F99
	cst13 = 0x24A19947B3916CF7
	cst14 = 0x0801F2E2858EFC16
	cst15 = 0x636920D871574E69
)

func block(d *digest, p []uint8) {
	h0, h1, h2, h3, h4, h5, h6, h7 := d.h[0], d.h[1], d.h[2], d.h[3], d.h[4], d.h[5], d.h[6], d.h[7]
	s0, s1, s2, s3 := d.s[0], d.s[1], d.s[2], d.s[3]

	for len(p) >= BlockSize {
		v0, v1, v2, v3, v4, v5, v6, v7 := h0, h1, h2, h3, h4, h5, h6, h7
		v8 := cst0 ^ s0
		v9 := cst1 ^ s1
		v10 := cst2 ^ s2
		v11 := cst3 ^ s3
		v12 := uint64(cst4)
		v13 := uint64(cst5)
		v14 := uint64(cst6)
		v15 := uint64(cst7)
		d.t += 1024
		if !d.nullt {
			v12 ^= d.t
			v13 ^= d.t
			v14 ^= 0 // TODO ideally d.t must be uint128.
			v15 ^= 0 //      currently 2^64 bits supported.
		}
		var m [16]uint64

		j := 0
		for i := 0; i < 16; i++ {
			m[i] = uint64(p[j])<<56 | uint64(p[j+1])<<48 | uint64(p[j+2])<<40 |
				uint64(p[j+3])<<32 | uint64(p[j+4])<<24 | uint64(p[j+5])<<16 |
				uint64(p[j+6])<<8 | uint64(p[j+7])
			j += 8
		}

		// Round 1.
		v0 += m[0] ^ cst1
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[2] ^ cst3
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[4] ^ cst5
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[6] ^ cst7
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[5] ^ cst4
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[7] ^ cst6
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[3] ^ cst2
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[1] ^ cst0
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[8] ^ cst9
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[10] ^ cst11
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[12] ^ cst13
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[14] ^ cst15
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[13] ^ cst12
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[15] ^ cst14
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[11] ^ cst10
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[9] ^ cst8
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 2.
		v0 += m[14] ^ cst10
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[4] ^ cst8
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[9] ^ cst15
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[13] ^ cst6
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[15] ^ cst9
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[6] ^ cst13
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[8] ^ cst4
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[10] ^ cst14
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[1] ^ cst12
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[0] ^ cst2
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[11] ^ cst7
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[5] ^ cst3
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[7] ^ cst11
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[3] ^ cst5
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[2] ^ cst0
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[12] ^ cst1
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 3.
		v0 += m[11] ^ cst8
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[12] ^ cst0
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[5] ^ cst2
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[15] ^ cst13
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[2] ^ cst5
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[13] ^ cst15
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[0] ^ cst12
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[8] ^ cst11
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[10] ^ cst14
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[3] ^ cst6
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[7] ^ cst1
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[9] ^ cst4
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[1] ^ cst7
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[4] ^ cst9
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[6] ^ cst3
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[14] ^ cst10
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 4.
		v0 += m[7] ^ cst9
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[3] ^ cst1
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[13] ^ cst12
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[11] ^ cst14
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[12] ^ cst13
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[14] ^ cst11
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[1] ^ cst3
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[9] ^ cst7
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[2] ^ cst6
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[5] ^ cst10
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[4] ^ cst0
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[15] ^ cst8
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[0] ^ cst4
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[8] ^ cst15
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[10] ^ cst5
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[6] ^ cst2
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 5.
		v0 += m[9] ^ cst0
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[5] ^ cst7
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[2] ^ cst4
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[10] ^ cst15
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[4] ^ cst2
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[15] ^ cst10
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[7] ^ cst5
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[0] ^ cst9
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[14] ^ cst1
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[11] ^ cst12
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[6] ^ cst8
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[3] ^ cst13
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[8] ^ cst6
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[13] ^ cst3
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[12] ^ cst11
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[1] ^ cst14
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 6.
		v0 += m[2] ^ cst12
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[6] ^ cst10
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[0] ^ cst11
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[8] ^ cst3
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[11] ^ cst0
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[3] ^ cst8
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[10] ^ cst6
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[12] ^ cst2
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[4] ^ cst13
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[7] ^ cst5
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[15] ^ cst14
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[1] ^ cst9
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[14] ^ cst15
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[9] ^ cst1
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[5] ^ cst7
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[13] ^ cst4
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 7.
		v0 += m[12] ^ cst5
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[1] ^ cst15
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[14] ^ cst13
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[4] ^ cst10
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[13] ^ cst14
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[10] ^ cst4
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[15] ^ cst1
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[5] ^ cst12
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[0] ^ cst7
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[6] ^ cst3
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[9] ^ cst2
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[8] ^ cst11
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[2] ^ cst9
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[11] ^ cst8
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[3] ^ cst6
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[7] ^ cst0
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 8.
		v0 += m[13] ^ cst11
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[7] ^ cst14
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[12] ^ cst1
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[3] ^ cst9
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[1] ^ cst12
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[9] ^ cst3
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[14] ^ cst7
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[11] ^ cst13
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[5] ^ cst0
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[15] ^ cst4
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[8] ^ cst6
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[2] ^ cst10
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[6] ^ cst8
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[10] ^ cst2
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[4] ^ cst15
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[0] ^ cst5
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 9.
		v0 += m[6] ^ cst15
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[14] ^ cst9
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[11] ^ cst3
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[0] ^ cst8
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[3] ^ cst11
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[8] ^ cst0
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[9] ^ cst14
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[15] ^ cst6
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[12] ^ cst2
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[13] ^ cst7
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[1] ^ cst4
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[10] ^ cst5
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[4] ^ cst1
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[5] ^ cst10
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[7] ^ cst13
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[2] ^ cst12
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 10.
		v0 += m[10] ^ cst2
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[8] ^ cst4
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[7] ^ cst6
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[1] ^ cst5
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[6] ^ cst7
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[5] ^ cst1
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[4] ^ cst8
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[2] ^ cst10
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[15] ^ cst11
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[9] ^ cst14
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[3] ^ cst12
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[13] ^ cst0
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[12] ^ cst3
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[0] ^ cst13
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[14] ^ cst9
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[11] ^ cst15
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 11.
		v0 += m[0] ^ cst1
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[2] ^ cst3
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[4] ^ cst5
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[6] ^ cst7
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[5] ^ cst4
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[7] ^ cst6
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[3] ^ cst2
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[1] ^ cst0
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[8] ^ cst9
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[10] ^ cst11
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[12] ^ cst13
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[14] ^ cst15
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[13] ^ cst12
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[15] ^ cst14
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[11] ^ cst10
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[9] ^ cst8
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 12.
		v0 += m[14] ^ cst10
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[4] ^ cst8
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[9] ^ cst15
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[13] ^ cst6
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[15] ^ cst9
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[6] ^ cst13
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[8] ^ cst4
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[10] ^ cst14
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[1] ^ cst12
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[0] ^ cst2
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[11] ^ cst7
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[5] ^ cst3
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[7] ^ cst11
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[3] ^ cst5
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[2] ^ cst0
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[12] ^ cst1
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 13.
		v0 += m[11] ^ cst8
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[12] ^ cst0
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[5] ^ cst2
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[15] ^ cst13
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[2] ^ cst5
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[13] ^ cst15
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[0] ^ cst12
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[8] ^ cst11
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[10] ^ cst14
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[3] ^ cst6
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[7] ^ cst1
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[9] ^ cst4
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[1] ^ cst7
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[4] ^ cst9
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[6] ^ cst3
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[14] ^ cst10
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 14.
		v0 += m[7] ^ cst9
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[3] ^ cst1
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[13] ^ cst12
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[11] ^ cst14
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[12] ^ cst13
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[14] ^ cst11
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[1] ^ cst3
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[9] ^ cst7
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[2] ^ cst6
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[5] ^ cst10
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[4] ^ cst0
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[15] ^ cst8
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[0] ^ cst4
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[8] ^ cst15
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[10] ^ cst5
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[6] ^ cst2
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 15.
		v0 += m[9] ^ cst0
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[5] ^ cst7
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[2] ^ cst4
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[10] ^ cst15
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[4] ^ cst2
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[15] ^ cst10
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[7] ^ cst5
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[0] ^ cst9
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[14] ^ cst1
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[11] ^ cst12
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[6] ^ cst8
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[3] ^ cst13
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[8] ^ cst6
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[13] ^ cst3
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[12] ^ cst11
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[1] ^ cst14
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		// Round 16.
		v0 += m[2] ^ cst12
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-32) | v12>>32
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-25) | v4>>25
		v1 += m[6] ^ cst10
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-32) | v13>>32
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-25) | v5>>25
		v2 += m[0] ^ cst11
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-32) | v14>>32
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-25) | v6>>25
		v3 += m[8] ^ cst3
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-32) | v15>>32
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-25) | v7>>25
		v2 += m[11] ^ cst0
		v2 += v6
		v14 ^= v2
		v14 = v14<<(64-16) | v14>>16
		v10 += v14
		v6 ^= v10
		v6 = v6<<(64-11) | v6>>11
		v3 += m[3] ^ cst8
		v3 += v7
		v15 ^= v3
		v15 = v15<<(64-16) | v15>>16
		v11 += v15
		v7 ^= v11
		v7 = v7<<(64-11) | v7>>11
		v1 += m[10] ^ cst6
		v1 += v5
		v13 ^= v1
		v13 = v13<<(64-16) | v13>>16
		v9 += v13
		v5 ^= v9
		v5 = v5<<(64-11) | v5>>11
		v0 += m[12] ^ cst2
		v0 += v4
		v12 ^= v0
		v12 = v12<<(64-16) | v12>>16
		v8 += v12
		v4 ^= v8
		v4 = v4<<(64-11) | v4>>11
		v0 += m[4] ^ cst13
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-32) | v15>>32
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-25) | v5>>25
		v1 += m[7] ^ cst5
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-32) | v12>>32
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-25) | v6>>25
		v2 += m[15] ^ cst14
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-32) | v13>>32
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-25) | v7>>25
		v3 += m[1] ^ cst9
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-32) | v14>>32
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-25) | v4>>25
		v2 += m[14] ^ cst15
		v2 += v7
		v13 ^= v2
		v13 = v13<<(64-16) | v13>>16
		v8 += v13
		v7 ^= v8
		v7 = v7<<(64-11) | v7>>11
		v3 += m[9] ^ cst1
		v3 += v4
		v14 ^= v3
		v14 = v14<<(64-16) | v14>>16
		v9 += v14
		v4 ^= v9
		v4 = v4<<(64-11) | v4>>11
		v1 += m[5] ^ cst7
		v1 += v6
		v12 ^= v1
		v12 = v12<<(64-16) | v12>>16
		v11 += v12
		v6 ^= v11
		v6 = v6<<(64-11) | v6>>11
		v0 += m[13] ^ cst4
		v0 += v5
		v15 ^= v0
		v15 = v15<<(64-16) | v15>>16
		v10 += v15
		v5 ^= v10
		v5 = v5<<(64-11) | v5>>11

		h0 ^= v0 ^ v8 ^ s0
		h1 ^= v1 ^ v9 ^ s1
		h2 ^= v2 ^ v10 ^ s2
		h3 ^= v3 ^ v11 ^ s3
		h4 ^= v4 ^ v12 ^ s0
		h5 ^= v5 ^ v13 ^ s1
		h6 ^= v6 ^ v14 ^ s2
		h7 ^= v7 ^ v15 ^ s3

		p = p[BlockSize:]
	}
	d.h[0], d.h[1], d.h[2], d.h[3], d.h[4], d.h[5], d.h[6], d.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}

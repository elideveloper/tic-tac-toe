package board

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func uintFromBitString(str string) uint8 {
	row, err := strconv.ParseUint(str, 2, 8)
	if err != nil {
		panic(err)
	}
	return uint8(row)
}

func TestIsWin(t *testing.T) {
	isBot := false
	testCases := []struct {
		name        string
		boardPlayer BoardMatrix
		needResp    bool
	}{
		{
			name: "true, first line horizontal",
			boardPlayer: BoardMatrix{
				uintFromBitString("00011111"),
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("00000000"),
				uintFromBitString("00010000"),
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("00000000"),
			},
			needResp: true,
		},
		{
			name: "true, last line horizontal",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("00000000"),
				uintFromBitString("00010000"),
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("00000000"),
				uintFromBitString("11111000"),
			},
			needResp: true,
		},
		{
			name: "true, middle line horizontal",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("10000001"),
				uintFromBitString("01000001"),
				uintFromBitString("00000000"),
				uintFromBitString("00111110"),
				uintFromBitString("00010000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, first column vertical",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000001"),
				uintFromBitString("11000001"),
				uintFromBitString("10000110"),
				uintFromBitString("10000000"),
				uintFromBitString("10000000"),
				uintFromBitString("00010000"),
				uintFromBitString("00010001"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, last column vertical",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000001"),
				uintFromBitString("11000001"),
				uintFromBitString("00000111"),
				uintFromBitString("10010001"),
				uintFromBitString("00010001"),
				uintFromBitString("00000000"),
				uintFromBitString("00010001"),
				uintFromBitString("00000001"),
			},
			needResp: true,
		},
		{
			name: "true, middle column vertical",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000000"),
				uintFromBitString("11000000"),
				uintFromBitString("00000111"),
				uintFromBitString("00010001"),
				uintFromBitString("00010000"),
				uintFromBitString("00010000"),
				uintFromBitString("00010001"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, up right diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000000"),
				uintFromBitString("11000000"),
				uintFromBitString("00100111"),
				uintFromBitString("00010001"),
				uintFromBitString("00011000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000001"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, middle right diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("00000000"),
				uintFromBitString("11000000"),
				uintFromBitString("00100101"),
				uintFromBitString("00010001"),
				uintFromBitString("00011000"),
				uintFromBitString("00000100"),
				uintFromBitString("00000000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, down right diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("00000000"),
				uintFromBitString("11000000"),
				uintFromBitString("00000101"),
				uintFromBitString("00010001"),
				uintFromBitString("00011000"),
				uintFromBitString("00000100"),
				uintFromBitString("00000010"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, up left diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000001"),
				uintFromBitString("10000010"),
				uintFromBitString("00000111"),
				uintFromBitString("00011001"),
				uintFromBitString("00010000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000001"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, middle left diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000101"),
				uintFromBitString("00011001"),
				uintFromBitString("00010000"),
				uintFromBitString("00100000"),
				uintFromBitString("01000000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, down left diag",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000001"),
				uintFromBitString("00011001"),
				uintFromBitString("00010000"),
				uintFromBitString("00100000"),
				uintFromBitString("01000000"),
				uintFromBitString("10010001"),
			},
			needResp: true,
		},

		{
			name: "true, shifted left diag in head",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000010"),
				uintFromBitString("00000100"),
				uintFromBitString("00001001"),
				uintFromBitString("00010001"),
				uintFromBitString("00110000"),
				uintFromBitString("00100000"),
				uintFromBitString("00000000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, shifted left diag in middle",
			boardPlayer: BoardMatrix{
				uintFromBitString("10000000"),
				uintFromBitString("00000100"),
				uintFromBitString("00001001"),
				uintFromBitString("00010001"),
				uintFromBitString("00100000"),
				uintFromBitString("01100000"),
				uintFromBitString("00000000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, shifted left diag in bottom",
			boardPlayer: BoardMatrix{
				uintFromBitString("10001000"),
				uintFromBitString("00010100"),
				uintFromBitString("00101000"),
				uintFromBitString("01000001"),
				uintFromBitString("10100000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000000"),
				uintFromBitString("00010001"),
			},
			needResp: true,
		},
		{
			name: "true, shifted right diag in head",
			boardPlayer: BoardMatrix{
				uintFromBitString("00100000"),
				uintFromBitString("00010000"),
				uintFromBitString("00001001"),
				uintFromBitString("00010101"),
				uintFromBitString("00100010"),
				uintFromBitString("00000000"),
				uintFromBitString("00000000"),
				uintFromBitString("00010000"),
			},
			needResp: true,
		},
		{
			name: "true, shifted right diag in middle",
			boardPlayer: BoardMatrix{
				uintFromBitString("00000000"),
				uintFromBitString("00010000"),
				uintFromBitString("00001001"),
				uintFromBitString("00010101"),
				uintFromBitString("00100010"),
				uintFromBitString("00000001"),
				uintFromBitString("00000000"),
				uintFromBitString("00010000"),
			},
			needResp: true,
		},
		{
			name: "true, shifted right diag in bottom",
			boardPlayer: BoardMatrix{
				uintFromBitString("00000000"),
				uintFromBitString("00000000"),
				uintFromBitString("00000001"),
				uintFromBitString("01010101"),
				uintFromBitString("00100000"),
				uintFromBitString("00010001"),
				uintFromBitString("00001000"),
				uintFromBitString("00010100"),
			},
			needResp: true,
		},
		{
			name: "false, random",
			boardPlayer: BoardMatrix{
				uintFromBitString("00100000"),
				uintFromBitString("01000100"),
				uintFromBitString("00010001"),
				uintFromBitString("00010101"),
				uintFromBitString("00100000"),
				uintFromBitString("00010101"),
				uintFromBitString("00001001"),
				uintFromBitString("10010000"),
			},
			needResp: false,
		},
		{
			name: "false, 4-max",
			boardPlayer: BoardMatrix{
				uintFromBitString("01111000"),
				uintFromBitString("01000101"),
				uintFromBitString("00011001"),
				uintFromBitString("01010101"),
				uintFromBitString("00101001"),
				uintFromBitString("10000100"),
				uintFromBitString("00001010"),
				uintFromBitString("00010100"),
			},
			needResp: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := Board{
				playerMoves: tc.boardPlayer,
			}
			resp := b.IsWin(isBot)
			assert.Equal(t, tc.needResp, resp)
		})
	}
}

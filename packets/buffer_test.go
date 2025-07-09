package packets

import (
	"bytes"
	"testing"
	"unicode/utf16"
)

func TestBuffer_WriteS(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []byte{0, 0},
		},
		{
			name:  "ASCII characters",
			input: "Hello",
			expected: []byte{
				'H', 0, 'e', 0, 'l', 0, 'l', 0, 'o', 0,
				0, 0,
			},
		},
		{
			name:  "Russian characters",
			input: "Привет",
			expected: []byte{
				0x1f, 0x04, 0x40, 0x04, 0x38, 0x04, 0x32, 0x04, 0x35, 0x04, 0x42, 0x04,
				0, 0,
			},
		},
		{
			name:  "Mixed ASCII and Unicode",
			input: "Hello 世界!",
			expected: []byte{
				'H', 0, 'e', 0, 'l', 0, 'l', 0, 'o', 0, ' ', 0,
				0x16, 0x4e, 0x4c, 0x75,
				'!', 0,
				0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buf := &Buffer{}

			buf.WriteS(tt.input)

			if !bytes.Equal(buf.b, tt.expected) {
				t.Errorf("\nОжидалось: %v\nПолучено:  %v", formatBytes(tt.expected), formatBytes(buf.b))
			}

			decoded := decodeUtf16LEWithTerminator(buf.b)
			if decoded != tt.input {
				t.Errorf("Декодированная строка не совпадает:\nОжидалось: %q\nПолучено:  %q", tt.input, decoded)
			}
		})
	}
}

func formatBytes(data []byte) string {
	var result string
	for i, b := range data {
		if i > 0 {
			result += " "
		}
		result += string([]byte{b})
	}
	return result
}

func decodeUtf16LEWithTerminator(data []byte) string {
	if len(data) < 2 || data[len(data)-2] != 0 || data[len(data)-1] != 0 {
		return ""
	}
	data = data[:len(data)-2]

	u16s := make([]uint16, len(data)/2)
	for i := 0; i < len(u16s); i++ {
		u16s[i] = uint16(data[i*2]) | uint16(data[i*2+1])<<8
	}

	return string(utf16.Decode(u16s))
}

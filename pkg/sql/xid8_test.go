package sql

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXID8_Scan(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected XID8
		wantErr  bool
		errMsg   string
	}{
		// Nil input
		{
			name:    "nil input",
			input:   nil,
			wantErr: true,
			errMsg:  "cannot scan nil value into XID8",
		},

		// Integer types (pgx style)
		{
			name:     "int64 positive",
			input:    int64(12345),
			expected: XID8(12345),
		},
		{
			name:     "int64 zero",
			input:    int64(0),
			expected: XID8(0),
		},
		{
			name:     "int64 max safe value",
			input:    int64(math.MaxInt64),
			expected: XID8(math.MaxInt64),
		},
		{
			name:    "int64 negative",
			input:   int64(-1),
			wantErr: true,
			errMsg:  "cannot convert negative int64",
		},
		{
			name:    "int64 large negative",
			input:   int64(-9223372036854775808), // math.MinInt64
			wantErr: true,
			errMsg:  "cannot convert negative int64",
		},

		// uint64 types
		{
			name:     "uint64 positive",
			input:    uint64(12345),
			expected: XID8(12345),
		},
		{
			name:     "uint64 zero",
			input:    uint64(0),
			expected: XID8(0),
		},
		{
			name:     "uint64 max value",
			input:    uint64(math.MaxUint64),
			expected: XID8(math.MaxUint64),
		},

		// String types
		{
			name:     "string positive number",
			input:    "12345",
			expected: XID8(12345),
		},
		{
			name:     "string zero",
			input:    "0",
			expected: XID8(0),
		},
		{
			name:     "string large number",
			input:    "18446744073709551615", // math.MaxUint64
			expected: XID8(math.MaxUint64),
		},
		{
			name:     "string with leading zeros",
			input:    "00012345",
			expected: XID8(12345),
		},
		{
			name:     "empty string",
			input:    "",
			expected: XID8(0),
		},
		{
			name:    "string with negative number",
			input:   "-12345",
			wantErr: true,
			errMsg:  "cannot parse string",
		},
		{
			name:    "string too large for uint64",
			input:   "18446744073709551616", // MaxUint64 + 1
			wantErr: true,
			errMsg:  "cannot parse string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var x XID8
			err := x.Scan(tt.input)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, x)
			}
		})
	}
}

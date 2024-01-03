package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"regexp:\\d+|len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: App{
				Version: "1",
			},
			expectedErr: fmt.Errorf(
				"field: Version | err: tagStringValidate len error - tag validation case mismatch\n", //nolint:revive
			),
		},
		{
			in: App{
				Version: "12345",
			},
			expectedErr: fmt.Errorf(""),
		},
		{
			in: Response{
				Body: "anytext",
				Code: 503,
			},
			expectedErr: fmt.Errorf(
				"field: Code | err: tagIntValidate in error - tag validation case subset len mismatch\n"), //nolint:revive
		},
		{
			in: Token{
				Header:    []byte{1, 2, 3},
				Payload:   []byte{4, 5, 6},
				Signature: []byte{7, 8, 9},
			},
			expectedErr: fmt.Errorf(""),
		},
		{
			in: User{
				ID:     "lessthan36symbols",
				Name:   "torchun",
				Age:    36,
				Email:  "some@example.com",
				Role:   "admin",
				Phones: []string{"12345", "1234567890", "0"}, // len != 11
				meta:   nil,
			},
			expectedErr: fmt.Errorf(
				"field: ID | err: tagStringValidate len error - tag validation case mismatch\n" +
					"field: Phones | err: tagStringValidate len error - tag validation case mismatch\n" +
					"field: Phones | err: tagStringValidate len error - tag validation case mismatch\n" +
					"field: Phones | err: tagStringValidate len error - tag validation case mismatch\n"),
		},
		{
			in: User{
				ID:     "exactly_36_symbols_long_abcdefghijkl",
				Name:   "torchun",
				Age:    36,
				Email:  "some@example.com",
				Role:   "admin",
				Phones: []string{"12345678901", "89995553311"},
				meta:   nil,
			},
			expectedErr: fmt.Errorf(""),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			require.EqualError(t, Validate(tt.in), tt.expectedErr.Error())
		})
	}
}

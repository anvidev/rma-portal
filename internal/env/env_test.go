package env

import (
	"os"
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name string
		k    string
		f    string
		v    string
		set  bool
		want string
	}{
		{
			name: "Get a variable that is set",
			k:    "TEST_SET_KEY",
			f:    "fallback",
			v:    "123abc",
			set:  true,
			want: "123abc",
		},
		{
			name: "Get a variable that is empty",
			k:    "TEST_EMPTY_KEY",
			f:    "fallback",
			v:    "",
			set:  true,
			want: "fallback",
		},
		{
			name: "Get a variable that is not set",
			k:    "TEST_UNSET_KEY",
			f:    "fallback",
			v:    "",
			set:  false,
			want: "fallback",
		}, {
			name: "Get a variable with non-ASCII characters",
			k:    "TEST_NON_ASCII_KEY",
			f:    "fallback",
			v:    "こんにちは",
			set:  true,
			want: "こんにちは",
		}, {
			name: "Get a variable with leading/trailing whitespace",
			k:    "TEST_WHITESPACE_KEY",
			f:    "fallback",
			v:    "  value  ",
			set:  true,
			want: "  value  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				os.Setenv(tt.k, tt.want)
				t.Cleanup(func() { os.Unsetenv(tt.k) })
			}
			got := GetString(tt.k, tt.f)
			if got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	tests := []struct {
		name string
		k    string
		f    time.Duration
		v    string
		set  bool
		want time.Duration
	}{
		{
			name: "Get a variable that is set to 3 hours (3h)",
			k:    "TEST_SET_KEY",
			f:    time.Hour * 24,
			v:    "3h",
			set:  true,
			want: time.Hour * 3,
		},
		{
			name: "Get a variable that is empty",
			k:    "TEST_EMPTY_KEY",
			f:    time.Hour * 24,
			v:    "",
			set:  true,
			want: time.Hour * 24,
		},
		{
			name: "Get a variable that is not set",
			k:    "TEST_UNSET_KEY",
			f:    time.Microsecond * 250,
			v:    "",
			set:  false,
			want: time.Microsecond * 250,
		},
		{
			name: "Get a variable that has incorrect value",
			k:    "TEST_INCORRECT_KEY",
			f:    time.Minute * 250,
			v:    "500minutes",
			set:  true,
			want: time.Minute * 250,
		}, {
			name: "Get a variable with negative duration",
			k:    "TEST_NEGATIVE_KEY",
			f:    time.Hour,
			v:    "-10s",
			set:  true,
			want: -10 * time.Second,
		},
		{
			name: "Get a variable with invalid duration format",
			k:    "TEST_INVALID_KEY",
			f:    time.Minute,
			v:    "10xyz",
			set:  true,
			want: time.Minute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				os.Setenv(tt.k, tt.v)
				t.Cleanup(func() { os.Unsetenv(tt.k) })
			}
			got := GetDuration(tt.k, tt.f)
			if got != tt.want {
				t.Errorf("GetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name string
		k    string
		f    int
		v    string
		set  bool
		want int
	}{
		{
			name: "Get a variable that is set to 100",
			k:    "TEST_SET_KEY",
			f:    250,
			v:    "100",
			set:  true,
			want: 100,
		},
		{
			name: "Get a variable that is empty",
			k:    "TEST_EMPTY_KEY",
			f:    100,
			v:    "",
			set:  true,
			want: 100,
		},
		{
			name: "Get a variable that is not set",
			k:    "TEST_UNSET_KEY",
			f:    100,
			v:    "",
			set:  false,
			want: 100,
		}, {
			name: "Get a variable with negative integer",
			k:    "TEST_NEGATIVE_KEY",
			f:    100,
			v:    "-50",
			set:  true,
			want: -50,
		},
		{
			name: "Get a variable with non-integer value",
			k:    "TEST_NON_INTEGER_KEY",
			f:    100,
			v:    "abc",
			set:  true,
			want: 100,
		},
		{
			name: "Get a variable with very large integer",
			k:    "TEST_LARGE_INT_KEY",
			f:    100,
			v:    "999999999999999999",
			set:  true,
			want: 999999999999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set {
				os.Setenv(tt.k, tt.v)
				t.Cleanup(func() { os.Unsetenv(tt.k) })
			}
			got := GetInt(tt.k, tt.f)
			if got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

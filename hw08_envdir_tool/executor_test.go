package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("env test: empty & unset", func(t *testing.T) {
		mapEnv := make(map[string]EnvValue)

		mapEnv["EMPTY"] = EnvValue{
			Value:      "",
			NeedRemove: false,
		}

		mapEnv["UNSET"] = EnvValue{
			Value:      "",
			NeedRemove: true,
		}

		code := RunCmd([]string{"echo", "test test test"}, mapEnv)

		_, okEmpty := os.LookupEnv("EMPTY")
		_, okUnset := os.LookupEnv("UNSET")

		require.Equal(t, code, 0)
		require.Equal(t, okEmpty, true)
		require.Equal(t, okUnset, false)
	})
}

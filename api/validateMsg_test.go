package api

import (
	"fmt"
	"testing"
	"time"
)

func TestValidateMsg(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		val            string
		expectedOutput string
	}{
		{
			val:            "This is a kerfuffle opinion I need to share with the world",
			expectedOutput: "This is a **** opinion I need to share with the world",
		},
		{
			val:            "This is a sharbert opinion I need to share with the world",
			expectedOutput: "This is a **** opinion I need to share with the world",
		},
		{
			val:            "This is a fornax opinion I need to share with the world",
			expectedOutput: "This is a **** opinion I need to share with the world",
		},
		{
			val:            "Hello fornax",
			expectedOutput: "Hello ****",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			parsedMsg := validateMsg(c.val)

			if parsedMsg != c.expectedOutput {
				t.Errorf(`Expected: %s Got: %s`, c.expectedOutput, parsedMsg)
				return
			}
		})
	}
}

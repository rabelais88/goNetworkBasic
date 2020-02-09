package statusCheck

import (
	"testing"
)

func TestCheck(t *testing.T) {
	res := Check([]string{`https://google.com`})
	if res == false {
		t.Error()
	}
}

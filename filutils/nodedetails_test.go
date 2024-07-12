package filutils

import (
	"fmt"
	"testing"
)

func TestNodeDetails(t *testing.T) {

	details := NodeDetails("f02824533")
	fmt.Println(details)
}

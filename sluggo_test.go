package sluggo

import (
	"fmt"
	"testing"
)

func TestSluggo(t *testing.T) {
	tStr := "Hello 123@email.com!"
	fmt.Println(Sluggo(tStr) == "hello-123-email-com")

	if Sluggo(tStr) != "hello-123-email-com" {
		t.Fatalf("Sluggo(\"%s\") should be \"hello-123-email-com\" but is %s", tStr, Sluggo(tStr))
	}
}

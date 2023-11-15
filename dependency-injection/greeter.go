package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(b *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

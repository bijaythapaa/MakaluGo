package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

func OtherGreet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return
	}
}

func main() {
	OtherGreet(os.Stdout, "Vijay")
}

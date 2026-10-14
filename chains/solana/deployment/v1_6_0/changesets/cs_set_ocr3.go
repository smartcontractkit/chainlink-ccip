package changesets

import (
	"fmt"
)

func runSafely(ops ...func()) {
	for _, op := range ops {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovered from panic: %v\n", r)
				}
			}()
			op()
		}()
	}
}

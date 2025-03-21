package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func SafeConvertWithBounds(n int) (uint, error) {
	if n < 0 {
		return 0, fmt.Errorf("out of range conversion: %d", n)
	}
	return uint(n), nil
}

func ExitOnError(err error, logger *zap.SugaredLogger) {
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
}

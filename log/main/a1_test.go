package log

import (
	"fmt"
	"testing"

	logger "github.com/bjk543/golib/log"
)

func TestGetArticles2(t *testing.T) {
	logger.Log("INFO", fmt.Sprintf("time %d", 123))
}

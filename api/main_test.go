package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.setMode(gin.TestMode)
	os.Exit(m.Run())
}

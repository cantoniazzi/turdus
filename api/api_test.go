package api_test

import (
	"os"
	"testing"

	"github.com/cantoniazzi/turdus/api"
)

func TestMain(m *testing.M) {
	go api.Start()
	os.Exit(m.Run())
}

//go:build integration

package api_v1

import (
	"os"
	"testing"

	"github.com/NikolNikolaeva/project_weather/it/testbed"
)

func TestMain(m *testing.M) {
	os.Exit(testbed.Bootstrap(m))
}

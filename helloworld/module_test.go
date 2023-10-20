package helloworld

import (
	"testing"

	"github.com/getsyncer/syncer-core/files"

	"github.com/getsyncer/syncer-core/drifttest"
)

func TestModule(t *testing.T) {
	config2 := `
version: 1
logic:
  - source: github.com/getsyncer/example-sync/helloworld
syncs:
  - logic: helloworld
    config:
      extra_content: "hello world"
`
	t.Run("with-config", drifttest.WithRun(config2, files.SimpleState(map[string]string{}), func(t *testing.T, items *drifttest.Items) {
		items.TestRun.MustExitCode(t, 0)
		drifttest.FileContains(t, "hello.txt", "hello world")
	}))
}

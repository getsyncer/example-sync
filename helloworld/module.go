package helloworld

import (
	"context"
	"fmt"

	"github.com/getsyncer/syncer-core/config"
	"github.com/getsyncer/syncer-core/drift"
	"github.com/getsyncer/syncer-core/drift/syncers/staticfile"
	"github.com/getsyncer/syncer-core/files"
	"github.com/getsyncer/syncer-core/fxregistry"
)

func init() {
	fxregistry.Register(Module)
}

const Name = config.Name("helloworld")

type Config struct {
}

func (c Config) Changes(ctx context.Context) (files.System[*files.StateWithChangeReason], error) {
	var ret files.System[*files.StateWithChangeReason]
	if err := ret.Add("hello.txt", &files.StateWithChangeReason{
		ChangeReason: &files.ChangeReason{
			Reason: "sync-output.txt",
		},
		State: files.State{
			Mode:          0644,
			Contents:      []byte(fmt.Sprintf("This file is synced by syncer-core!\n%s", drift.MagicTrackedString)),
			FileExistence: files.FileExistencePresent,
		},
	}); err != nil {
		return ret, err
	}
	return ret, nil
}

var Module = staticfile.NewCustomModule[Config](Name, drift.Priority(0))

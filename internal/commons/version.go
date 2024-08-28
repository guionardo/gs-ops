package commons

import (
	"time"

	"github.com/guionardo/gs-ops/internal/models/responses"
	"github.com/hashicorp/go-version"
)

const (
	AppName     = "gs-ops"
	VersionText = "0.0.1"
	VersionDate = "2024-08-27"
)

var Version responses.VersionResponse

func init() {
	var (
		err error
		v   *version.Version
		d   time.Time
	)
	if v, err = version.NewSemver(VersionText); err == nil {
		if d, err = time.Parse(time.DateOnly, VersionDate); err == nil {
			Version = responses.VersionResponse{
				AppName: AppName,
				Version: v.String(),
				Date:    d,
			}
			return
		}
	}
	panic(err)
}

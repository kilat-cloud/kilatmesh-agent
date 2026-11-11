package version

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// EnvAllowUpstreamUpdate is the explicit operator opt-in that re-enables the
// upstream self-update machinery (version checks and installer execution).
const EnvAllowUpstreamUpdate = "KILATMESH_ALLOW_UPSTREAM_UPDATE"

// UpstreamUpdateDisabledMessage explains why an update path refused to run.
const UpstreamUpdateDisabledMessage = "KilatMesh Agent updates are managed by PT Kilat Cloud Indonesia; upstream self-update is disabled in this build"

// UpstreamUpdateAllowed reports whether the upstream self-update machinery
// may run. KilatMesh Agent ships with upstream self-update disabled: a
// deployed agent must never replace itself with an upstream binary or poll
// upstream update infrastructure. Only an explicit operator opt-in via
// EnvAllowUpstreamUpdate turns it back on.
func UpstreamUpdateAllowed() bool {
	val := os.Getenv(EnvAllowUpstreamUpdate)
	if val == "" {
		return false
	}
	allowed, err := strconv.ParseBool(val)
	if err != nil {
		log.Warnf("failed to parse %s: %v", EnvAllowUpstreamUpdate, err)
		return false
	}
	return allowed
}

//go:build !cgo || !((linux && (amd64 || arm64)) || (darwin && arm64))

package ere

// ere publishes its verifier library only for linux/amd64, linux/arm64 and darwin/arm64, and linking it needs cgo.
// This undefined name makes any other build fail with a legible error instead of missing symbols.
var _ = ere_needs_cgo_on_linux_amd64_linux_arm64_or_darwin_arm64

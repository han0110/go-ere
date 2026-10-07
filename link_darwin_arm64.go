//go:build cgo

package ere

// #cgo LDFLAGS: ${SRCDIR}/lib/darwin_arm64/libere_verifier_c.a
import "C"

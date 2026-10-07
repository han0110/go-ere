//go:build cgo

package ere

// #cgo LDFLAGS: ${SRCDIR}/lib/linux_arm64/libere_verifier_c.a -lm -lpthread -ldl
import "C"

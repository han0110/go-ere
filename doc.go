// Package ere verifies zkVM proofs with the verifiers of github.com/eth-act/ere,
// through its C library ere-verifier-c.
//
// The library is prebuilt and committed under lib/ for each supported platform:
// linux/amd64, linux/arm64 and darwin/arm64. It is linked statically, so
// `go build` needs a C compiler and nothing else. On any other platform, or
// without cgo, the package does not compile.
//
// A *Verifier may be used from multiple goroutines for concurrent Verify and
// Kind calls. Close consumes the handle and must not overlap with other calls
// on the same Verifier.
package ere

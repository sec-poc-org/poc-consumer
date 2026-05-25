// Minimal program that no longer reaches any known-vulnerable function.
// Phase 3b PoC pre-state called jwt.MapClaims.VerifyAudience (GO-2020-0017
// reachable). PR B removes that call so govulncheck stops flagging it.
package main

import "fmt"

func main() {
	fmt.Println("clean program — no reachable vuln paths")
}

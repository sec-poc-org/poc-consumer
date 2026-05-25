// PoC for Phase 3b — govulncheck reachability scanning.
//
// Uses github.com/dgrijalva/jwt-go v3.2.0+incompatible.
//
// GO-2020-0017 (Authorization bypass in MapClaims.VerifyAudience) is the
// vulnerability of interest. This file calls VerifyAudience explicitly so the
// vulnerable function is REACHABLE from main — govulncheck will flag it as a
// Symbol Result (the strictest finding category, where the code actually
// exercises the buggy function), not just a package-level finding.
//
// To demonstrate reachability filtering: if you remove the VerifyAudience call
// below and keep the import, govulncheck will downgrade the finding to a
// Package Result and exit 0. That's the value-add over Dependabot, which
// flags any vulnerable version regardless of whether the buggy function is
// ever called.
package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	claims := jwt.MapClaims{"aud": "my-service"}

	// Reachable call to the vulnerable function.
	audValid := claims.VerifyAudience("my-service", true)
	fmt.Printf("audience valid: %v\n", audValid)
}

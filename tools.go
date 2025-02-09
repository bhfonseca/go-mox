package go_mox

import "fmt"

func PVEAuth(PVEUser, PVETokenID, PVEToken string) string {
	return fmt.Sprintf("PVEAPIToken=%s!%s=%s", PVEUser, PVETokenID, PVEToken)
}

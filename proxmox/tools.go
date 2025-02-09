package proxmox

import "fmt"

func PVEAuth(PVEUser, PVETokenID, PVEToken string) string {
	return fmt.Sprintf("PVEAPIToken=%s!%s=%s", PVEUser, PVETokenID, PVEToken)
}

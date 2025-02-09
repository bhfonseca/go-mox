package models

type PVENodes struct {
	Data []struct {
		Node      string `json:"node"`
		Status    string `json:"status"`
		Type      string `json:"type"`
		Uptime    int    `json:"uptime"`
		MaxDisk   int    `json:"maxdisk"`
		Disk      int    `json:"disk"`
		MaxCPU    int    `json:"maxcpu"`
		CPU       int    `json:"cpu"`
		MaxMemory int    `json:"maxmem"`
		Memory    int    `json:"mem"`
	} `json:"data"`
}

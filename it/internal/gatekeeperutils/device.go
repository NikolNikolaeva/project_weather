package gatekeeperutils

type Device struct {
	DeviceId   string `json:"deviceId,omitempty"`
	Name       string `json:"name,omitempty"`
	SiteName   string `json:"siteName,omitempty"`
	MacAddr    string `json:"macAddr,omitempty"`
	Notes      string `json:"notes,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

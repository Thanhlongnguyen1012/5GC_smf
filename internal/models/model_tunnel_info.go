package models

type TunnelInfo struct {
	Ipv4Addr string `json:"ipv4Addr,omitempty"`
	Ipv6Addr string `json:"ipv6Addr,omitempty"`
	GtpTeid  string `json:"gtpTeid"`
}

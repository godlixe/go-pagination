package entity

type Site struct {
	ID             uint64 `json:"id" gorm:"primaryKey"`
	GlobalRank     int64  `json:"global_rank"`
	TldRank        int64  `json:"tld_rank"`
	Domain         string `json:"domain"`
	TLD            string `json:"tld"`
	RefSubNets     string `json:"ref_sub_nets"`
	RefIPs         string `json:"ref_ips"`
	IDNDomain      string `json:"idn_domain"`
	IDNTld         string `json:"idn_tld"`
	PrevGlobalRank int64  `json:"prev_global_rank"`
	PrevTldRank    int64  `json:"prev_tld_rank"`
	PrevRefSubNets string `json:"prev_ref_sub_nets"`
	PrevRefIPs     string `json:"prev_ref_ips"`
}

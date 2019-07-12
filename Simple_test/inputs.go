package test

type Inputs struct {
	Vcn_display_name    string `json:"vcn_display_name"`
	Vcn_label           string `json:"vcn_label"`
	Vcn_cidr            string `json:"vcn_cidr"`
	Vcn_expectedSubnets int    `json:5`
}

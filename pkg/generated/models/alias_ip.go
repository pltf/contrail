package models

// AliasIP

// AliasIP
//proteus:generate
type AliasIP struct {
	UUID                 string              `json:"uuid,omitempty"`
	ParentUUID           string              `json:"parent_uuid,omitempty"`
	ParentType           string              `json:"parent_type,omitempty"`
	FQName               []string            `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType        `json:"id_perms,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs      `json:"annotations,omitempty"`
	Perms2               *PermType2          `json:"perms2,omitempty"`
	AliasIPAddress       IpAddressType       `json:"alias_ip_address,omitempty"`
	AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family,omitempty"`

	ProjectRefs                 []*AliasIPProjectRef                 `json:"project_refs,omitempty"`
	VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP {
	return &AliasIP{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		AliasIPAddress:       MakeIpAddressType(),
		AliasIPAddressFamily: MakeIpAddressFamilyType(),
	}
}

// MakeAliasIPSlice() makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
	return []*AliasIP{}
}

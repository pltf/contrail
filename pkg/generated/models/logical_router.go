package models

// LogicalRouter

// LogicalRouter
//proteus:generate
type LogicalRouter struct {
	UUID                      string           `json:"uuid,omitempty"`
	ParentUUID                string           `json:"parent_uuid,omitempty"`
	ParentType                string           `json:"parent_type,omitempty"`
	FQName                    []string         `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType     `json:"id_perms,omitempty"`
	DisplayName               string           `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs   `json:"annotations,omitempty"`
	Perms2                    *PermType2       `json:"perms2,omitempty"`
	VxlanNetworkIdentifier    string           `json:"vxlan_network_identifier,omitempty"`
	ConfiguredRouteTargetList *RouteTargetList `json:"configured_route_target_list,omitempty"`

	RouteTargetRefs             []*LogicalRouterRouteTargetRef             `json:"route_target_refs,omitempty"`
	VirtualMachineInterfaceRefs []*LogicalRouterVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceInstanceRefs         []*LogicalRouterServiceInstanceRef         `json:"service_instance_refs,omitempty"`
	RouteTableRefs              []*LogicalRouterRouteTableRef              `json:"route_table_refs,omitempty"`
	VirtualNetworkRefs          []*LogicalRouterVirtualNetworkRef          `json:"virtual_network_refs,omitempty"`
	PhysicalRouterRefs          []*LogicalRouterPhysicalRouterRef          `json:"physical_router_refs,omitempty"`
	BGPVPNRefs                  []*LogicalRouterBGPVPNRef                  `json:"bgpvpn_refs,omitempty"`
}

// LogicalRouterVirtualMachineInterfaceRef references each other
type LogicalRouterVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterServiceInstanceRef references each other
type LogicalRouterServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterRouteTableRef references each other
type LogicalRouterRouteTableRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterVirtualNetworkRef references each other
type LogicalRouterVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterPhysicalRouterRef references each other
type LogicalRouterPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterBGPVPNRef references each other
type LogicalRouterBGPVPNRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LogicalRouterRouteTargetRef references each other
type LogicalRouterRouteTargetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeLogicalRouter makes LogicalRouter
func MakeLogicalRouter() *LogicalRouter {
	return &LogicalRouter{
		//TODO(nati): Apply default
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		Perms2:                    MakePermType2(),
		VxlanNetworkIdentifier:    "",
		ConfiguredRouteTargetList: MakeRouteTargetList(),
	}
}

// MakeLogicalRouterSlice() makes a slice of LogicalRouter
func MakeLogicalRouterSlice() []*LogicalRouter {
	return []*LogicalRouter{}
}

// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

type (
	LogicalRouterStaticRoutePolicy = string
)

const (
	LogicalRouterStaticRoutePolicySrcIP LogicalRouterStaticRoutePolicy = "src-ip"
	LogicalRouterStaticRoutePolicyDstIP LogicalRouterStaticRoutePolicy = "dst-ip"
)

// LogicalRouterStaticRoute defines an object in Logical_Router_Static_Route table
type LogicalRouterStaticRoute struct {
	UUID        string                           `ovsdb:"_uuid"`
	ExternalIDs map[string]string                `ovsdb:"external_ids"`
	IPPrefix    string                           `ovsdb:"ip_prefix"`
	Nexthop     string                           `ovsdb:"nexthop"`
	Options     map[string]string                `ovsdb:"options"`
	OutputPort  []string                         `ovsdb:"output_port"`
	Policy      []LogicalRouterStaticRoutePolicy `ovsdb:"policy"`
}
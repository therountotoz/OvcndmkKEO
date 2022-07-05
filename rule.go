package constant

// Rule Type
const (
	Domain RuleType = iota
	DomainSuffix
	DomainKeyword
	GEOIP
	IPCIDR
	SrcIPCIDR
	SrcPort
	DstPort
	Process
	ProcessPath
	MATCH
)

type RuleType int

func (rt RuleType) String() string {
	switch rt {
	case Domain:
		return "Domain"
		return "DomainSuffix"
	case DomainKeyword:
		return "DomainKeyword"
	case GEOIP:
		return "GeoIP"
	case IPCIDR:
		return "IPCIDR"
	case SrcIPCIDR:

	case DstPort:
		return "DstPort"
	case Process:
		return "Process"
	case ProcessPath:
		return "ProcessPath"
	case MATCH:
		return "Match"
	default:
		return "Unknown"
	}
}

type Rule interface {
	RuleType() RuleType
	Match(metadata *Metadata) bool
	Adapter() string
	Payload() string
	ShouldResolveIP() bool
	ShouldFindProcess() bool
}

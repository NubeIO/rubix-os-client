package constants

type DeviceType int64

const (
	Cloud DeviceType = iota
	Edge28
	RubixCompute
	RubixComputeVPN
	RubixComputeLoRaWAN
	RubixComputeLoRaWANVPN
	RubixComputeIO
)

func ValidRubixCompute(s string) bool {
	switch s {
	case RubixCompute.String():
		return true
	case RubixComputeVPN.String():
		return true
	case RubixComputeLoRaWAN.String():
		return true
	case RubixComputeLoRaWANVPN.String():
		return true
	case RubixComputeIO.String():
		return true
	}
	return false
}

func (s DeviceType) String() string {
	switch s {
	case Cloud:
		return "cloud"
	case Edge28:
		return "edge-28"
	case RubixCompute:
		return "rubix-compute"
	case RubixComputeVPN:
		return "rubix-compute-vpn"
	case RubixComputeLoRaWAN:
		return "rubix-compute-lorawan"
	case RubixComputeLoRaWANVPN:
		return "rubix-compute--lorawan-vpn"
	case RubixComputeIO:
		return "rubix-compute-io"
	}
	return "unknown"
}

package dock

import "fmt"

type RequestType int

const (
	ReqInvalid             = -1
	ReqEnquire RequestType = iota
	ReqResetToBootloader
	ReqVersion
	ReqPower
	ReqName
	ReqDebug
	ReqSet
)

func (e RequestType) String() string {
	switch e {
	case ReqEnquire:
		return "Enquire"
	case ReqResetToBootloader:
		return "Reset To Bootloader"
	case ReqVersion:
		return "Version"
	case ReqPower:
		return "Power"
	case ReqName:
		return "Name"
	case ReqDebug:
		return "Debug"
	case ReqSet:
		return "Set"
	}
	return "invalid RequestType"
}

type Request struct {
	RequestType
	Channel  int
	Params   []int
	ParamStr string
}

func (e Request) String() string {
	if e.RequestType == ReqName {
		if e.Params[0] == int('u') {
			return fmt.Sprintf("Request: [%v user, %v]", e.RequestType, e.ParamStr)
		}
		if e.Params[0] == int('d') {
			return fmt.Sprintf("Request: [%v dock, %v]", e.RequestType, e.ParamStr)
		}
		return fmt.Sprintf("Request: [%v invalid(%v), %v]", e.RequestType, e.Params[0], e.ParamStr)
	}

	return fmt.Sprintf("Request: [%v, %v]", e.RequestType, e.Params)
}

package httpenc

type HTTPRespCode int

const (
	HTTPRespCode_Successful HTTPRespCode = iota + 1
	HTTPRespCode_ParamUnmarshal
	HTTPRespCode_ParamMiss
	HTTPRespCode_ParamIllegal
	HTTPRespCode_NotFound
	HTTPRespCode_Repeated
)

func HTTPRespCodeMessage(code HTTPRespCode) string {
	switch code {
	case HTTPRespCode_Successful:
		return "Successful"
	case HTTPRespCode_ParamUnmarshal:
		return "ParamUnmarshal"
	case HTTPRespCode_ParamMiss:
		return "ParamMiss"
	case HTTPRespCode_ParamIllegal:
		return "ParamIllegal"
	case HTTPRespCode_NotFound:
		return "NotFound"
	case HTTPRespCode_Repeated:
		return "Repeated"
	}
	return "Unknown"
}

type HTTPRespPackage struct {
	Code HTTPRespCode `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string       `json:"msg"`
}

func MakeSuccHTTPRespPackage(data interface{}) HTTPRespPackage {
	return HTTPRespPackage{
		Code: HTTPRespCode_Successful,
		Data: data,
	}
}

func MakeFailHTTPRespPackage(code HTTPRespCode) HTTPRespPackage {
	return HTTPRespPackage{
		Code: code,
		Msg:  HTTPRespCodeMessage(code),
	}
}

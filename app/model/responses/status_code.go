package responses

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 定義系統錯誤與相關回傳訊息
const (
	Success      = "0"
	ParameterErr = "1"
	Error        = "2"
)

var MsgText = map[string]string{
	Success:      "Success",
	ParameterErr: "Parameter error, please check your field.",
	Error:        "Has some promble",
}

func Status(code string, data interface{}) Response {
	return Response{
		Status:  code,
		Data:    data,
		Message: MsgText[code]}
}

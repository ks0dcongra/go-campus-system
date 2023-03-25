package responses

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Redis   string      `json:"redis"`
}

// 定義系統錯誤與相關回傳訊息
const (
	Success      = "0"
	ParameterErr = "1"
	Error        = "2"
	SuccessDb	 = "3"
	SuccessRedis = "4"
)

var MsgText = map[string]string{
	Success:      "Success",
	ParameterErr: "Parameter error, please check your field.",
	Error:        "Has some promble",
	SuccessDb:	  "Success from DB",
	SuccessRedis: "Success from Redis",
}

func Status(code string, data interface{}, redis string) Response {
	return Response{
		Status:  code,
		Data:    data,
		Message: MsgText[code],
		Redis:   redis,
	}
}

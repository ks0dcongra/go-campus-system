package responses

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// 定義系統錯誤與相關回傳訊息
const (
	Success       = "0"
	ParameterErr  = "1"
	Error         = "2"
	SuccessDb     = "3"
	SuccessRedis  = "4"
	ScoreTokenErr = "5"
	TokenErr      = "6"
	TokenExpired  = "7"
	PasswordErr   = "8"
	DbErr         = "9"
)

var MsgText = map[string]string{
	Success:       "Success",
	ParameterErr:  "Parameter error, please check your field.",
	Error:         "Has some problem",
	SuccessDb:     "Success from DB",
	SuccessRedis:  "Success from Redis",
	ScoreTokenErr: "The scores only limited myself to query",
	TokenErr:      "Token issue, can't pass authorization",
	TokenExpired:  "Token has been Expired because blaklist has this token",
	PasswordErr:   "Database hash password not in consonance with input password",
	DbErr:         "SQL not found",
}

func Status(code string, data interface{}) Response {
	return Response{
		Status:  code,
		Data:    data,
		Message: MsgText[code],
	}
}

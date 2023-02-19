package models

type Response struct {
	Data  any    `json:"data"`
	Msg   string `json:"msg"`
	Error any    `json:"error"`
}

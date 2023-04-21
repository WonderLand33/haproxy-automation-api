package server

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

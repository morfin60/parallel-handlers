package multiplexer

type Result struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Url     string `json:"url"`
	Content []byte `json:"content"`
}

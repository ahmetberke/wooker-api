package response

type AnyResponse struct {
	Response
	Data interface{} `json:"data"`
}
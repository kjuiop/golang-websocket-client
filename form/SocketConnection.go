package form

type SocketConnectionRequest struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	ApiPath      string `json:"api_path"`
	Protocol     string `json:"protocol"`
	ClientNumber int    `json:"client_number"`
	Period       int    `json:"period"`
}

type SocketConnectionResponse struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

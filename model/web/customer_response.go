package web

type CustomerResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

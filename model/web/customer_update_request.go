package web

type CustomerUpdateRequest struct {
	Id      int    `validate:"required" json:"id"`
	Name    string `validate:"required,max=200,min=1" json:"name"`
	Address string `validate:"required,max=200,min=1" json:"address"`
}

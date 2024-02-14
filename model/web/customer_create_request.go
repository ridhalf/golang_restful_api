package web

type CustomerCreateRequest struct {
	Name    string `validate:"required,max=200,min=1" json:"name"`
	Address string `validate:"required,max=200,min=1" json:"address"`
}

package request

type ProjectRequest struct {
	PublicId      string `json:"PublicId"`
	Title         string `json:"Title"`
	Keywords      string `json:"Keywords"`
	Domain        string `json:"Domain"`
	Description   string `json:"Description"`
	Address       string `json:"Address"`
	City          string `json:"City"`
	State         string `json:"State"`
	Zipcode       string `json:"Zipcode"`
	ProjectPhoto1 string `json:"ProjectPhoto1"`
	ProjectPhoto2 string `json:"ProjectPhoto2"`
	ProjectPhoto3 string `json:"ProjectPhoto3"`
	ProjectPhoto4 string `json:"ProjectPhoto4"`
}


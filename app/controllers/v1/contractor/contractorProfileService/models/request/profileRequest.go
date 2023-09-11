package request

type ContractorProfileRequest struct {
	ContractorId    int    `json:"Contractor_id"`
	PublicId        string `json:"publicId"`
	Title           string `json:"Title"`
	Domain          string `json:"Domain"`
	Bio             string `json:"Bio"`
	Phone           string `json:"Phone"`
	Address         string `json:"Address"`
	City            string `json:"City"`
	State           string `json:"State"`
	Zipcode         string `json:"Zipcode"`
	Url             string `json:"Url"`
	ProfilePhoto    string `json:"ProfilePhoto"`
	BackgroundPhoto string `json:"BackgroundPhoto"`
	TwitterHandle   string `json:"TwitterHandle"`
	FacebookHandle  string `json:"FacebookHandle"`
	InstagramHandle string `json:"InstagramHandle"`
	LinkedinHandle  string `json:"LinkedinHandle"`
}

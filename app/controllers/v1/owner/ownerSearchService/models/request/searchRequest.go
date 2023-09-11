package request

type SearchRequest struct {
	Domain  string `json:"Domain"`
	Zipcode string `json:"Zipcode"`
}

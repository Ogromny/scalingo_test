package datatypes

type ReposResponseBody struct {
	Items []ReposResponseBodyItems `json:"items"`
}

type ReposResponseBodyItems struct {
	FullName string `json:"full_name"`
	Language string `json:"language"`
	Url      string `json:"html_url"`
}

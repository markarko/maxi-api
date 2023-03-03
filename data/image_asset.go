package data

type ImageAsset struct {
	ImageUrl        string `json:"imageUrl"`
	SmallUrl        string `json:"smallUrl"`
	MediumUrl       string `json:"mediumUrl"`
	LargeUrl        string `json:"largeUrl"`
	SmallRetinaUrl  string `json:"smallRetinaUrl"`
	MediumRetinaUrl string `json:"mediumRetinaUrl"`
	LargeRetinaUrl  string `json:"largeRetinaUrl"`
}

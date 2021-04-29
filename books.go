package main

type Root struct {
	Value []Books `json:"value"`
}

type Books struct {
	RowId       int
	ImageUrl    string
	ISBN        string
	Price       string
	PublishDate string
	Title       string
	URL         string
}

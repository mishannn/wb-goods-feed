package feed

type Post struct {
	Title   string
	Content string
	Images  []Image
	Link    string
}

type Image struct {
	URL string
}

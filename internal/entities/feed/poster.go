package feed

type Poster interface {
	PublishPost(post Post) error
}

package feed

type URLShortener interface {
	GetShortURL(longURL string) (string, error)
}

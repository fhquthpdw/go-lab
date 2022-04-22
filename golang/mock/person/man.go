package person

//go:generate mockgen -destination=../mock/black_mock.go -package=mock mock/person Black

type Man interface {
    Walk(road string) string
    Talk(sentence string) string
}

package data

type Data_access interface {
	Save(key string, value string)
	Load(key string) string
}

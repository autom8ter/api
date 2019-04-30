package driver

type Database interface {
	Save(Document) error
	Update(Document) error
	Get(category string, document string) (Document, error)
}

type Document interface {
	DocCategory() string
	DocName() string
	DocData() map[string]interface{}
}

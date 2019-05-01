package driver

type Map map[string]string

func (m Map) Meta() map[string]string {
	return m
}

func AsMeta(m map[string]string) Metadata {
	return Map(m)
}

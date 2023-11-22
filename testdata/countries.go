package countries

type Country struct {
	name string
}

var TR = Country{
	name: "a",
}

func (c Country) Name() string {
	return c.name
}
func AllCountries() []Country {
	return []Country{
		TR,
	}
}

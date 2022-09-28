package internal

type DB struct {
	Examples []Example
}

type Example struct {
	Description string
	Command     string
	Content     string
}

func (d *DB) AddExample(e Example) {
	if e.Command == "" && e.Description == "" {
		return
	}
	d.Examples = append(d.Examples, e)
}

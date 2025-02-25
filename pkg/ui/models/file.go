package models

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type File struct {
	Name     string
	Size     int64
	Modified string
}

func (f *File) Render() Node {
	return Tr(
		Td(Text(f.Name)),
		Td(Text(fmt.Sprint(f.Size))),
		Td(Text(f.Modified)),
	)
}

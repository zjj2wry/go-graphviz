package dot

import "strings"

type Attrs struct {
	Items []string
}

func (o *Attrs) AddAttrs(attrs string) {
	o.Items = append(o.Items, strings.Split(attrs, " ")...)
}

func (o *Attrs) String() string {
	return strings.Join(o.Items, " ")
}

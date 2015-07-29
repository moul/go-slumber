package slumber

import "strings"

type Slumber struct {
	Base string
}

func (s *Slumber) Child(resource string) *Slumber {
	return &Slumber{
		Base: strings.TrimRight(s.Base, "/") + "/" + resource,
	}
}

func New(base string) *Slumber {
	return &Slumber{
		Base: base,
	}
}

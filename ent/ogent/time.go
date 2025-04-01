package ogent

import (
	"time"
)

type Datetime time.Time

func (d *Datetime) UnmarshalParam(param string) error {
	t, err := time.Parse(`2006-01-02 15:04:05 MST`, param)
	if err != nil {
		return err
	}
	*d = Datetime(t)
	return nil
}

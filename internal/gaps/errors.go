package gaps

import "fmt"

var (
	ErrGoldenNotFound  = newErr("golden profile not found")
	ErrProfileNotFound = newErr("default profile not found")
)

type Err struct {
	msg, meta string
}

func newErr(msg string) *Err {
	return &Err{msg: msg}
}

func (e *Err) WithMeta(meta string) *Err {
	if meta != "" {
		e.meta = meta
	}

	return e
}

func (e *Err) Error() string {
	if e.meta != "" {
		return fmt.Sprintf("%s\n%s", e.msg, e.meta)
	}
	return e.msg
}

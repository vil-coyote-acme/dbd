package storage

// TODO add some details

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

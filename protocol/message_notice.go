package protocol

// error
type Error struct {
	Message string
}

type Chat struct {
	Pos     int
	Message string
}

type Announce struct {
	Message string
}

func NewError(err error) *Message {
	return NewMessage(Error{
		Message: err.Error(),
	})
}

func NewChat(body string) *Message {
	return NewMessage(Chat{
		Message: body,
	})
}

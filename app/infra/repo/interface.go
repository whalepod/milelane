package repo

type NoteAccessor interface {
	Create(title string, body string) error
}

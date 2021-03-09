package repo

// NoteAccessor gives access to repo layer from handler.
type NoteAccessor interface {
	Create(title string, body string) error
}

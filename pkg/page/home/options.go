package home

type Option struct {
	title  string
	keyTag rune
}

func (o Option) isQuit() bool {
	return o.keyTag == 'q'
}

var options = []Option{
	{title: "✏️  - [A]dd a new note", keyTag: 'a'},
	{title: "📒 - [V]iew all notes", keyTag: 'v'},
	{title: "📝 - [E]dit a note", keyTag: 'e'},
	{title: "🗑️  - [D]elete a note", keyTag: 'd'},
	{title: "🔍 - [S]earch notes by keyword", keyTag: 's'},
	{title: "🚪 - [Q]uit", keyTag: 'q'},
}

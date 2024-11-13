package mnc

type History struct {
	Messages []*Message	`json:"messages"`
}

func (m Message) String() string {
	return m.Content
}

func (h *History) Save(message *Message) {
	h.Messages = append(h.Messages, message)
}

func (h History) List() string {
	list := ""
	for _, m := range h.Messages {
		list += m.Content
	}
	return list
}

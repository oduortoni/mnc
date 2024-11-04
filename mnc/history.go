package mnc

type History struct {
	Messages []string
}

func (h *History) Save(message string) {
	h.Messages = append(h.Messages, message)
}

func (h History) List() string {
	list := ""
	for _, m := range h.Messages {
		list += m
	}
	return list
}

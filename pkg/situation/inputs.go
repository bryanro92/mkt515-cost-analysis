package situation

type answers struct {
	ans []string
}

type questions struct {
	q []string
}

func (q *questions) populateQ() {
	q.q = []string{
		"Enter initial cost:",
		"Enter total target market (number of customers / year):",
	}
}

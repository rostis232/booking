package forms

type errors map[string][]string

// Add adds an error message for the given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	} else {
		return es[0]
	}
}

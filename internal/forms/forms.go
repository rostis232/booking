package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds an url.Values object
type Form struct {
	url.Values
	Errors errors
}


// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	} else {
		return true
	}
}

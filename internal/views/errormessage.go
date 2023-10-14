package views

type ErrorMessageView struct {
	ErrorMessage string `json:"message"`
}

func ErrorMessageViewFrom(errMsg string) ErrorMessageView {
	view := ErrorMessageView{
		ErrorMessage: errMsg,
	}
	return view
}

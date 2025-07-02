package helper


func HandleErrors(err error) *int {
	var code int
	if err == nil {
		return  nil
	}
	switch err.Error() {
		case BADREQUEST:
			code = 400
		case INTERNALSERVER:
			code = 500
		case NOTAUTHORIZED:
			code = 401
		case NOTFOUND:
			code = 404
	}
	return  &code
}

func SetProperty(status int) string {
	if status >= 400 {
		return "error"
	} else {
		return  "message"
	}
}
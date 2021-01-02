package log

type Log struct {
	code          int
	log_message   string
	error_message error
	detail        string
}

type LogMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type ErrorMessage struct {
	Code   int          `json:"code"`
	Error  string       `json:"error"`
	Trace  []ErrorTrace `json:"trace"`
	Detail string       `json:"detail"`
}

type ErrorTrace struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (l *Log) SetCode(_code_ int) {
	l.code = _code_
}

func (l *Log) SetLogMessage(_log_message_ string) {
	l.log_message = _log_message_
}

func (l *Log) SetErrorMessage(_error_message_ error) {
	l.error_message = _error_message_
}

func (l *Log) SetDetail(_detail_ string) {
	l.detail = _detail_
}

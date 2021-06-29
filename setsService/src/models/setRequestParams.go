package models

type SetRequestParamsUnParsed struct {
	IsPublic string `omitempty`
	IsMine   string `omitempty`
	Limit    string `omitempty`
	Offset   string `omitempty`
	SetName  string `omitempty`
}

type SetRequestParams struct {
	IsPublic bool   `omitempty`
	IsMine   bool   `omitempty`
	Limit    int    `omitempty`
	Offset   int    `omitempty`
	SetName  string `omitempty`
}

type SetRequestParamsError struct {
	ErrorMessage string `json:"errorMessage"`
}

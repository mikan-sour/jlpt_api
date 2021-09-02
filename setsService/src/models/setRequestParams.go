package models

type SetRequestById struct {
	Id string `omitempty`
}

type SetRequestParamsUnParsed struct {
	Id       string `omitempty`
	IsPublic string `omitempty`
	IsMine   string `omitempty`
	Limit    string `omitempty`
	Offset   string `omitempty`
	SetName  string `omitempty`
}

type SetRequestParams struct {
	Id       string `omitempty`
	IsPublic bool   `omitempty`
	IsMine   bool   `omitempty`
	Limit    int    `omitempty`
	Offset   int    `omitempty`
	SetName  string `omitempty`
}

type SetRequestParamsError struct {
	ErrorMessage string `json:"errorMessage"`
}

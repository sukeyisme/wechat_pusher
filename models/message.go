package models

type Message struct {
	ToUser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	Url        string `json:"url"`
	Data       Data   `json:"data"`
}

type Data struct {
	First    Raw `json:"first"`
	Keyword1 Raw `json:"keyword1"`
	Keyword2 Raw `json:"keyword2"`
	Keyword3 Raw `json:"keyword3"`
	Remark   Raw `json:"remark"`
}

type Raw struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

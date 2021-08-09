package json_structs

type UserPostRequestBody struct {
	LineToken                string `json:"line_token" validate:"required"`
	IsCook                   bool   `json:"is_cook"`
	GetResponseNotifications bool   `json:"get_response_notifications"`
	Password                 string `json:"password" validate:"required,min=8,max=30"`
}

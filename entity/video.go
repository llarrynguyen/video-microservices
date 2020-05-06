package entity

type Person struct {
	FirstName string `json: "firstname" binding:"required"`
	LastName  string `json: "lastName" binding:"required"`
	Age       int8   `json: "age" binding: "gte=1, lte=30"`
	Email     string `json: "email" validate: "required,email"`
}

type Video struct {
	Title       string `json: "title" binding:"min=2,max=500" validate="is-cool"`
	Description string `json: "description binding:"min=2,max=500"`
	URL         string `json: "url" binding: "required,url"`
	Author      Person `json:"author" binding:"required"`
}

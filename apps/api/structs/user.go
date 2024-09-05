package structs

type User struct {
	ID       string `json:"id" pg:"id,notnull,pk,type:'uuid'"`
	Email    string `json:"email" pg:"email,notnull,unique"`
	Password string `json:"password" pg:"password,notnull"`
}

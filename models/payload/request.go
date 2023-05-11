package payload

type Login struct {
	Username string `json:"Username" form:"Username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Register struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type CreateCategory struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type CreateTask struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Status      string `json:"status" form:"status"`
	CategoryID  uint   `json:"category_id" form:"category_id" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

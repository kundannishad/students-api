package student

type StudentStr struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
	Age   int    `json:"age,omitempty" validate:"required"`
}

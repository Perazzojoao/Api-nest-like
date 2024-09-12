package entity

type UserEntity struct {
	Name string
	Age  int
}

func NewUserEntity() *UserEntity {
	name := "João"
	age := 25
	return &UserEntity{
		Name: name,
		Age:  age,
	}
}

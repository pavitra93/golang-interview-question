package builder

type student struct {
	name  string
	email string
	phone string
}

func NewStudent(name string, email string, phone string) *student {
	return &student{name: name, email: email, phone: phone}
}

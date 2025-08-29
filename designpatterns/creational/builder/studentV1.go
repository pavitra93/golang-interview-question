package builder

type Student struct {
	name  string
	email string
	phone string
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetEmail(email string) {
	s.email = email
}

func (s *Student) SetPhone(phone string) {
	s.phone = phone
}

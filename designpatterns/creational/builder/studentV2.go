package builder

type studentV2 struct {
	name     string
	email    string
	phone    string
	isActive bool
}

type StudentV2Builder struct {
	studentV2 *studentV2
}

func GetInstance() *StudentV2Builder {
	return &StudentV2Builder{studentV2: &studentV2{}}
}

func (builder *StudentV2Builder) SetName(name string) *StudentV2Builder {
	builder.studentV2.name = name
	return builder
}

func (builder *StudentV2Builder) SetEmail(email string) *StudentV2Builder {
	builder.studentV2.email = email
	return builder
}

func (builder *StudentV2Builder) SetPhone(phone string) *StudentV2Builder {
	builder.studentV2.phone = phone
	return builder
}

func (builder *StudentV2Builder) SetActive() *StudentV2Builder {
	builder.studentV2.isActive = true
	return builder
}

func (builder *StudentV2Builder) Build() *studentV2 {
	return &studentV2{
		name:     builder.studentV2.name,
		email:    builder.studentV2.email,
		phone:    builder.studentV2.phone,
		isActive: builder.studentV2.isActive,
	}
}

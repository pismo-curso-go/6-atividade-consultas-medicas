package patients

type PatientDomain struct {
	id       int
	name     string
	email    string
	password string
}

func NewPatientDomain(
	name string,
	email string,
	password string,
) *PatientDomain {
	return &PatientDomain{
		name:     name,
		email:    email,
		password: password,
	}
}

func NewPatientDomainFromDB(id int, name, email, password string) *PatientDomain {
	return &PatientDomain{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}

func (p *PatientDomain) ID() int {
	return p.id
}

func (p *PatientDomain) Name() string {
	return p.name
}

func (p *PatientDomain) Email() string {
	return p.email
}

func (p *PatientDomain) Password() string {
	return p.password
}

func (p *PatientDomain) Validate() error {
	if p.name == "" {
		return ErrPatientBlankName
	}

	if p.email == "" {
		return ErrPatientBlankEmail
	}

	if p.password == "" {
		return ErrPatientBlankPassword
	}

	return nil
}

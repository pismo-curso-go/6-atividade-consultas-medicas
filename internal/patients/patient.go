package patients

type PatientDomain struct {
	name     string
	email    string
	password string
}

func NewPatientDomain(
	name string,
	email string,
	password string,
) (*PatientDomain, error) {
	if name == "" {
		return nil, ErrPatientInvalidName
	}

	if email == "" {
		return nil, ErrPatientInvalidEmail
	}

	return &PatientDomain{
		name:     name,
		email:    email,
		password: password,
	}, nil
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

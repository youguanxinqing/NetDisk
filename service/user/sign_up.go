package user

type SignUpService struct {
	Username string
	Password string
	Tel      string
}

func (srv *SignUpService) Rigister() error {
	return nil
}

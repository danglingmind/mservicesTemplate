package auth

type AuthInterface interface {
}

type AuthModel struct {
}

var _ AuthInterface = &AuthModel{}

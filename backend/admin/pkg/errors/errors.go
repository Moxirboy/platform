package errors

const (
	ErrUserNotFound       = Err("user not found")
	ErrEmptyUserName      = Err("user name can not be empty")
	ErrEmptyPhoneNumber   = Err("phone number can not be empty")
	ErrInvalidCredentials = Err("bad credentials")
	ErrPhoneNumberExists  = Err("this phone number already exists")
	ErrInvalidPassword    = Err("invalid password")
	ErrClassNotFound      = Err("class not found")
)

type Err string

func (e Err) Error() string {
	return string(e)
}

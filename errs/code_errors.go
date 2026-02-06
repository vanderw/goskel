package errs

// CodeError Compatible with Error but additionally with error-code.
type CodeError struct {
	s string
	c int
}

func New(s string, c int) error {
	return &CodeError{
		s: s,
		c: c,
	}
}

func (ce *CodeError) Error() string {
	return ce.s
}

func (ce *CodeError) Code() int {
	return ce.c
}

var (
	ErrParams       = New("invalid parameters", -1)
	ErrNotEmpty     = New("not empty data", -2)
	ErrEmpty        = New("empty data", -3)
	ErrFormat       = New("format error", -5)
	ErrNotEqual     = New("value not equal", -7)
	ErrEqual        = New("value is equal", -8)
	ErrInsufficient = New("length or cap is insufficient", -9)
	ErrOutOfRange   = New("index out of range", -10)
	ErrKeyMissing   = New("key not exists", -51)
	ErrKeyExists    = New("key exists already", -61)
	ErrParsing      = New("data parsing failed", -70)
)

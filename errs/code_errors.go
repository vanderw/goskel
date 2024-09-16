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
	ErrEmpty        = New("empty data", -2)
	ErrFormat       = New("format error", -3)
	ErrNotEqual     = New("value not equal", -5)
	ErrEqual        = New("value is equal", -6)
	ErrInsufficient = New("length or cap is insufficient", -7)
	ErrOutOfRange   = New("index out of ramge", -7)
	ErrKeyMissing   = New("key not exists", -51)
	ErrKeyExists    = New("key exists already", -61)
	ErrParsing      = New("data parsing failed", -70)
)

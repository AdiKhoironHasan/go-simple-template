package consts

type ContextKey string

const (
	CtxRequestId ContextKey = "request_id"
	CtxUser      ContextKey = "user"
)

func (c ContextKey) String() string {
	return string(c)
}

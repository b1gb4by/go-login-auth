package util

type ErrorCode int

const (
	ErrorCode00001 ErrorCode = iota + 1
	ErrorCode00002
	ErrorCode00003
	ErrorCode00004
	ErrorCode00005
	ErrorCode00006
)

const (
	ErrorCode10000 ErrorCode = iota + 10000
	ErrorCode10001
	ErrorCode10002
	ErrorCode10003
	ErrorCode10004
	ErrorCode10005
	ErrorCode10006
	ErrorCode10007
	ErrorCode10008
	ErrorCode10009
	ErrorCode10010
)

const (
	UnknownError ErrorCode = 99999
)

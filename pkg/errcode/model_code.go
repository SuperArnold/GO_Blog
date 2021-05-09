package errcode

var (
	ErrorCreateTagFail = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail = NewError(20010003, "更新標籤失敗")
)

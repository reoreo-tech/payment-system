// ginフレームワークにアプリケーションをすべて依存させないように
// interfaceを使ってgin.COntextを抽象化する
package handler

// Context - context interface
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}

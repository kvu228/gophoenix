package engine

/*
Interface IrouterMap có 2 abstract functions
InsertRoute(path, handler) gắn handler vào path tương ứng
SearchRoute(path) trả về handler và path tương ứng
*/
type IRouterMap interface {
	InsertRoute(path string, handler HandlerFunc)
	SearchRoute(path string) (HandlerFunc, map[string]any)
}

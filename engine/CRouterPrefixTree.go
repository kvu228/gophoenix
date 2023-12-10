// CRouterPrefixTree kế thừa IRouter

package engine

import (
	"fmt"
	"strings"
)

type PrefixNode struct {
	childs         []*PrefixNode
	key            string
	isDynamicParam bool
	handler        HandlerFunc
	fullKeyPath    string
}

// Xử lý path -> trả về array gồm các key của path
func splitPath(path string) []string {
	trimPath := strings.TrimPrefix(path, "/")
	arr := strings.Split(trimPath, "/")
	return arr
}

// trả về PrefixNode tương ướng với key
func (n *PrefixNode) matchChildWithKey(key string) *PrefixNode {
	for _, child := range n.childs {
		if child.key == key || child.isDynamicParam {
			return child
		}
	}
	return nil

}

// Khai báo kiểu CRouterPrefixTree
type CRouterPrefixTree struct {
	root *PrefixNode
}

// Cấp phát mới (giống _init_ bên python)
func NewCRouterPrefixTree() *CRouterPrefixTree {
	return &CRouterPrefixTree{
		root: &PrefixNode{},
	}
}

// Thực thi InsertRoute
func (c *CRouterPrefixTree) InsertRoute(path string, handler HandlerFunc) {
	keys_path := splitPath(path)
	fmt.Println("Insert Keys Path: ", keys_path, len(keys_path))

	//nếu tại "/" trả về handler của root
	if path == "/" {
		c.root.handler = handler
		return
	}


	curNode := c.root
	for _, key := range keys_path {
		childNode := curNode.matchChildWithKey(key)
		if childNode == nil {
			childNode = &PrefixNode{}
			childNode.key = key
			childNode.isDynamicParam = string(key[0]) == ":"
			curNode.childs = append(curNode.childs, childNode)
		}
		curNode = childNode
	}

	curNode.handler = handler
	curNode.fullKeyPath = path
	fmt.Println(curNode.key, curNode.handler)
}

//                ""
//	              api
//             :version
//     :id_user    :id_product
// user               product


// search: /api/v2/34/product
//Thực thi SearchRoute
func (c *CRouterPrefixTree) SearchRoute(runtimePath string) (HandlerFunc, map[string]any) {

	key_paths := splitPath(runtimePath)
	fmt.Println("search key paths: ", key_paths)
	searchedNode := c.searchLevelRouter(c.root, key_paths, 0)
	if searchedNode == nil {
		return nil, nil
	}

	savedPaths := splitPath(searchedNode.fullKeyPath)
	params := make(map[string]any)
	for index, savedPath := range savedPaths {
		keyPath := key_paths[index]
		if keyPath != savedPath {
			params[savedPath[1:]] = keyPath
			fmt.Println(savedPath[1:], keyPath)
		}
	}

	fmt.Println(params)
	return searchedNode.handler, params
}


func (c *CRouterPrefixTree) searchLevelRouter(curNode *PrefixNode, path []string, indexLevel int) *PrefixNode {
	curKey := path[indexLevel]
	if curKey == "" && curNode.handler != nil {
		return curNode
	}
	for _, childNode := range curNode.childs {
		if childNode.key == curKey || childNode.isDynamicParam {
			if childNode == nil {
				return nil
			}
			if childNode != nil {
				if childNode.handler != nil && indexLevel == len(path)-1 {
					return childNode
				} else {
					// tam dung level o day va di sau xuong con cua no
					return c.searchLevelRouter(childNode, path, indexLevel+1)
				}
			}
		}
	}
	return nil
}

// body cua levelRouter cua node api level 1 van o day
// body cua levelrouter cua rootNode dau tien van o day

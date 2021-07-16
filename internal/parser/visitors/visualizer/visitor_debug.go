package visualizer

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/arborlang/ArborGo/internal/parser/ast"
	"github.com/arborlang/ArborGo/internal/parser/visitors/base"
)

type TreeInfo struct {
	Name     string      `json:"name"`
	Parent   *string     `json:"parent"`
	Children []*TreeInfo `json:"children"`
}

func (t *TreeInfo) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

type Visualizer struct {
	treeInfo *TreeInfo
	stacks   []*TreeInfo
}

func (v *Visualizer) SetVisitor(bv *base.VisitorAdapter) {

}

func str(value string) *string {
	return &value
}

func (v *Visualizer) getCurrentLevel() *TreeInfo {
	return v.stacks[len(v.stacks)-1]
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (v *Visualizer) appendToCurrentLevel(n ast.Node) {
	currentScope := v.getCurrentLevel()
	currentScope.Children = append(currentScope.Children, &TreeInfo{
		Name:     fmt.Sprintf("%s (%s)", getType(n), n.GetType().String()),
		Children: []*TreeInfo{},
	})
}

func (v *Visualizer) appendName(msg string, values ...interface{}) *TreeInfo {
	tree := &TreeInfo{
		Name:     fmt.Sprintf(msg, values...),
		Children: []*TreeInfo{},
	}
	if v.treeInfo == nil {
		v.treeInfo = tree
		v.stacks = append(v.stacks, v.treeInfo)
		return v.treeInfo
	}
	currentScope := v.getCurrentLevel()
	currentScope.Children = append(currentScope.Children, tree)
	return tree
}

func (v *Visualizer) makeNewLevel(t *TreeInfo) {
	v.stacks = append(v.stacks, t)
}

func (v *Visualizer) popLevel() *TreeInfo {
	i := len(v.stacks) - 1 // Any valid index, however you happen to get it.
	x := v.stacks[i]
	v.stacks = v.stacks[:i]
	return x
}

func New() (*Visualizer, ast.Visitor) {
	v := &Visualizer{}
	v.stacks = []*TreeInfo{v.treeInfo}
	return v, v
}

func handler(v *Visualizer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("webpage").Parse(tmpl)
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, struct{ Value string }{Value: v.treeInfo.String()})
		if err != nil {
			log.Println(err)
		}
	}
}

func NewServer(v *Visualizer) error {
	log.Println("Starting visualizer server on port 8080")
	http.HandleFunc("/", handler(v))
	return http.ListenAndServe(":8080", nil)
}

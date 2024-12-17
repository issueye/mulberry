package stand

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type GojaModule struct {
	name string
	sets map[string]interface{}

	runtime *goja.Runtime
	once    sync.Once
}

func NewGojaModule(name string) Module {
	return &GojaModule{
		name: name,
		sets: make(map[string]interface{}),
	}
}

func (p *GojaModule) String() string {
	return p.name
}

func (p *GojaModule) Name() string {
	return p.name
}

func (p *GojaModule) Set(objects Objects) Module {
	for k, v := range objects {
		p.sets[k] = v
	}

	return p
}

func (p *GojaModule) Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)

	for k, v := range p.sets {
		o.Set(k, v)
	}
}

func (p *GojaModule) Enable(runtime Runtime) {
	runtime.Set(p.Name(), require.Require(runtime.(*goja.Runtime), p.Name()))
}

func (p *GojaModule) Register() Module {
	p.once.Do(func() {
		require.RegisterNativeModule(p.Name(), p.Require)
	})

	return p
}

type Objects map[string]interface{}

type Runtime interface {
	Set(name string, value interface{}) error
}

type Object interface {
	Set(string, interface{})
	Get(string) interface{}
}

type Module interface {
	Name() string
	Set(objects Objects) Module
	Enable(Runtime)
	Register() Module
}

type TemplateVars struct {
	PackageName  string
	PackagePath  string
	PackageVars  map[string]string
	PackageTypes map[string]string
	PackageFuncs map[string]string

	Args map[string]interface{}
}

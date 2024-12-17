package code_engine

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/require"
)

const (
	LogKey   = "log"
	DebugKey = "debug"
	InfoKey  = "info"
	ErrorKey = "error"
	WarnKey  = "warn"
	SelfKey  = "self"
	AtobKey  = "atob"
	BtoaKey  = "btoa"
)

type Code struct {
	Path    string
	Program *goja.Program
}

type JsVM struct {
	// vm 虚拟机
	vm *goja.Runtime
	// 注册
	registry *require.Registry
	// 全局goja加载目录
	globalPath string
	// 输出回调
	ConsoleCallBack ConsoleCallBack
	// 外部添加到内部的内容
	pkg map[string]map[string]any
	// 对应文件的编译对象
	proMap  map[string]*Code
	console *console
}

type ModuleFunc = func(vm *goja.Runtime, module *goja.Object)

func NewJsVM(globalPath string, console *console, consoleCallBack ConsoleCallBack) *JsVM {
	jsVM := &JsVM{
		vm:              goja.New(),
		globalPath:      globalPath,
		ConsoleCallBack: consoleCallBack,
		pkg:             make(map[string]map[string]any),
		proMap:          make(map[string]*Code),
	}

	jsVM.vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	// 输出日志
	o := jsVM.vm.NewObject()
	o.Set(LogKey, console.Log)
	o.Set(DebugKey, console.Debug)
	o.Set(InfoKey, console.Info)
	o.Set(ErrorKey, console.Error)
	o.Set(WarnKey, console.Warn)
	jsVM.vm.Set("console", o)

	jsVM.console = console
	if consoleCallBack != nil {
		console.CallBack = append(console.CallBack, &consoleCallBack)
	}

	var parserOpts []parser.Option
	jsVM.vm.SetParserOptions(parserOpts...)

	ops := []require.Option{}

	if globalPath != "" {
		ops = append(ops, require.WithGlobalFolders(globalPath))
	}

	// source
	// sourceLoader := jsVM.sourceLoader(globalPath)
	// ops = append(ops, require.WithLoader(sourceLoader))

	jsVM.registry = require.NewRegistry(ops...)
	jsVM.registry.Enable(jsVM.vm)

	self := jsVM.vm.GlobalObject()
	jsVM.vm.Set(SelfKey, self)

	jsVM.vm.Set(AtobKey, func(code string) string {
		raw, err := Base64DecodeString(code)
		if err != nil {
			panic(err)
		}
		return raw
	})

	jsVM.vm.Set(BtoaKey, func(code string) string {
		return Base64EncodeString(code)
	})

	return jsVM
}

func (jv *JsVM) Set(name string, value any) {
	jv.vm.Set(name, value)
}

func (jv *JsVM) load() {
	// 加载其他模块
	for name, mod := range jv.pkg {
		gojaMod := jv.vm.NewObject()
		for k, v := range mod {
			gojaMod.Set(k, v)
		}

		// 注册模块
		jv.vm.Set(name, gojaMod)
	}
}

// SetProperty
// 向模块写入变量或者写入方法
func (jv *JsVM) SetProperty(moduleName, key string, value any) {
	mod, ok := jv.pkg[moduleName]
	if !ok {
		jv.pkg[moduleName] = make(map[string]any)
		mod = jv.pkg[moduleName]
	}
	mod[key] = value
}

func (jv *JsVM) SetConsoleCallBack(consoleCallBack ConsoleCallBack) {
	jv.console.SetCallBack(&consoleCallBack)
}

func (jv *JsVM) Run(name string, pro *goja.Program) error {
	// 加载模块
	jv.load()

	if pro != nil {
		var exception error
		_, err := jv.vm.RunProgram(pro)
		if err != nil {
			gojaErr, ok := err.(*goja.Exception)
			if !ok {
				exception = err
			} else {
				exception = errors.New(gojaErr.String())
			}
		}

		if exception != nil {
			return exception
		}
	} else {
		return errors.New("code is nil")
	}

	return nil
}

func (jv *JsVM) compile(name string, path string) (pro *goja.Program, err error) {
	var tmpPath string
	if jv.globalPath != "" {
		tmpPath = filepath.Join(jv.globalPath, path)
	} else {
		tmpPath = path
	}

	// 规范化路径
	tmpPath = filepath.Clean(tmpPath)

	// 读取文件
	src, err := os.ReadFile(tmpPath)
	if err != nil {
		return nil, err
	}

	// 编译文件
	pro, err = goja.Compile(name, string(src), false)
	if err != nil {
		fmt.Println("compile error:", err)
		return nil, err
	}

	jv.proMap[name] = &Code{
		Path:    path,
		Program: pro,
	}

	return
}

func (jv *JsVM) wrapCode(name, code string) string {
	return fmt.Sprintf(`
    function %s() {
            const module = (function(exports) {
                %s
                return exports
            })(self)
            return module.%s();
    }`, name, code, name)
}

func (jv *JsVM) RunCode(code string) error {
	_, err := jv.vm.RunString(code)
	return err
}

func (jv *JsVM) ExportFunc(name string, path string, fn any) error {
	// 输入验证
	if name == "" || path == "" {
		return fmt.Errorf("invalid input: name and path must not be empty")
	}

	// 编译
	pro, err := jv.compile(name, path)
	if err != nil {
		return err
	}

	vm := jv.vm
	// 运行
	err = jv.Run(name, pro)
	if err != nil {
		return err
	}

	nameFunc := vm.Get(name)
	if nameFunc == goja.Undefined() || nameFunc == goja.Null() {
		return fmt.Errorf("%s function not found", name)
	}
	_, ok := goja.AssertFunction(nameFunc)
	if !ok {
		return fmt.Errorf("%s function not found", name)
	}

	// 导出
	return vm.ExportTo(vm.Get(name), fn)
}

package codegen

import (
	"fmt"
	"log"
	"strings"

	"github.com/progrium/macdriver/internal/set"

	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

// Function is code generator for objective-c function
type Function struct {
	Type        *typing.FunctionType
	Name        string // the first part of objc function name
	GoName      string
	Parameters  []*Param
	ReturnType  typing.Type
	Deprecated  bool // if has been deprecated
	Suffix      bool // GoName conflicts so add suffix to this function
	Description string
	DocURL      string

	goFuncName string
	identifier string
}

// GoArgs return go function args
func (f *Function) GoArgs(currentModule *modules.Module) string {
	// log.Println("rendering function", f.Name)
	var args []string
	for _, p := range f.Parameters {
		log.Println("rendering function", f.Name, p.Name, p.Type)
		log.Printf("rendering function ptype: %T", p.Type)
		if pt, ok := p.Type.(*typing.PointerType); ok {
			log.Printf("ptr type: %T", pt.Type)
		}
		args = append(args, fmt.Sprintf("%s %s", p.Name, p.Type.GoName(currentModule, true)))
	}
	return strings.Join(args, ", ")
}

// GoReturn return go function return
func (f *Function) GoReturn(currentModule *modules.Module) string {
	if f.ReturnType == nil {
		return ""
	}
	return f.ReturnType.GoName(currentModule, true)
}

// CArgs return go function args
func (f *Function) CArgs(currentModule *modules.Module) string {
	// log.Println("rendering function", f.Name)
	var args []string
	for _, p := range f.Parameters {
		args = append(args, fmt.Sprintf("%s %s", p.Type.ObjcName(), p.Name))
	}
	return strings.Join(args, ", ")
}

// Selector return full Objc function name
func (f *Function) Selector() string {
	if f.identifier == "" {
		var sb strings.Builder
		sb.WriteString(f.Name)
		for idx, p := range f.Parameters {
			if idx > 0 {
				sb.WriteString(p.FieldName)
			}
			sb.WriteString(":")
		}
		f.identifier = sb.String()
	}
	return f.identifier
}

func (f *Function) String() string {
	return f.Selector()
}

// NormalizeInstanceTypeFunction return new init function.
func (f *Function) NormalizeInstanceTypeFunction(returnType *typing.ClassType) *Function {
	nm := &Function{
		Name:       f.Name,
		GoName:     f.GoName,
		Parameters: f.Parameters,
		ReturnType: returnType,
		goFuncName: f.goFuncName,
		Suffix:     f.Suffix,
	}
	return nm
}

// WriteGoCallCode generate go function code to call c wrapper code
func (f *Function) WriteGoCallCode(currentModule *modules.Module, cw *CodeWriter) {
	funcDeclare := f.GoFuncDeclare(currentModule)

	if f.Deprecated {
		return
		cw.WriteLine("// deprecated")
	}

	if f.DocURL != "" {
		cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", f.Description))
		cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", f.DocURL))
	}

	cw.WriteLine("func " + funcDeclare + " {")
	cw.Indent()

	var returnTypeStr string
	rt := typing.UnwrapAlias(f.ReturnType)
	switch rt.(type) {
	case *typing.VoidType:
		returnTypeStr = "objc.Void"
	default:
		returnTypeStr = f.ReturnType.GoName(currentModule, true)
	}
	callCode := fmt.Sprintf("objc.Call[%s](%s, objc.Sel(\"%s\")", returnTypeStr, f.Selector())
	var sb strings.Builder
	for idx, p := range f.Parameters {
		sb.WriteString(", ")
		switch tt := p.Type.(type) {
		case *typing.ClassType:
			sb.WriteString("objc.Ptr(" + p.GoName() + ")")
		case *typing.ProtocolType:
			pvar := fmt.Sprintf("po%d", idx)
			cw.WriteLineF("%s := objc.WrapAsProtocol(\"%s\", %s)", pvar, tt.Name, p.GoName())
			sb.WriteString(pvar)
		case *typing.PointerType:
			switch tt.Type.(type) {
			case *typing.ClassType: //object pointer convert to unsafe.Pointer, avoiding ffi treat it as PointerHolder
				sb.WriteString(fmt.Sprintf("unsafe.Pointer(%s)", p.GoName()))
			default:
				sb.WriteString(p.GoName())
			}
		default:
			sb.WriteString(p.GoName())
		}
	}
	callCode += sb.String() + ")"

	switch rt.(type) {
	case *typing.VoidType:
		cw.WriteLine(callCode)
	default:
		var resultName = "rv"
		cw.WriteLine(resultName + " := " + callCode)
		cw.WriteLine("return " + resultName)
	}
	cw.UnIndent()
	cw.WriteLine("}")
}

func (f *Function) WriteCSignature(currentModule *modules.Module, cw *CodeWriter) {
	var returnTypeStr string
	rt := typing.UnwrapAlias(f.ReturnType)
	switch rt.(type) {
	case *typing.VoidType:
		returnTypeStr = "void *"
	default:
		returnTypeStr = f.ReturnType.ObjcName()
	}
	cw.WriteLineF("// %v %v(%v); ", returnTypeStr, f.GoName, f.CArgs(currentModule))
}

// WriteGoInterfaceCode generate go interface function signature code
func (f *Function) WriteGoInterfaceCode(currentModule *modules.Module, classType *typing.ClassType, w *CodeWriter) {
	if f.Deprecated {
		return
		w.WriteLine("// deprecated")
	}
	funcDeclare := f.GoFuncDeclare(currentModule)
	w.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (f *Function) GoFuncDeclare(currentModule *modules.Module) string {
	var returnType = f.GoReturn(currentModule)
	return f.Type.GoName(currentModule, true) + "(" + f.GoArgs(currentModule) + ") " + returnType
}

// GoImports return all imports for go file
func (f *Function) GoImports() set.Set[string] {
	var imports = set.New("github.com/progrium/macdriver/objc")
	for _, param := range f.Parameters {
		imports.AddSet(param.Type.GoImports())
	}
	if f.ReturnType != nil {
		imports.AddSet(f.ReturnType.GoImports())
	}
	return imports
}
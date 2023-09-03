package codegen

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/progrium/darwinkit/generate/modules"
	"github.com/progrium/darwinkit/generate/typing"
	"github.com/progrium/darwinkit/internal/set"
)

// ModuleWriter mantains module level auto-generated code source files
type ModuleWriter struct {
	Module      modules.Module
	Description string
	DocURL      string
	PlatformDir string
	Protocols   []*typing.ProtocolType
	EnumAliases []*AliasInfo
	FuncAliases []*AliasInfo
	Functions   []*Function
	Structs     []*Struct
}

func (m *ModuleWriter) WriteCode() {
	m.WriteDocFile()
	m.WriteEnumAliases()
	m.WriteTypeAliases()
	m.WriteStructs()
	m.WriteFunctions()
	m.WriteFunctionWrappers()
	if m.Module.Package == "coreimage" {
		// filter protocols maybe arent "real" protocols?
		// get "cannot find protocol declaration" with protocol imports
		return
	}
	if m.Module.Package == "coremediaio" {
		// get "cannot find protocol declaration" with protocol imports
		return
	}
	if m.Module.Package == "fileprovider" {
		// get "cannot find protocol declaration" with protocol imports
		return
	}
	m.WriteProtocolsImportCode()
}

func (m *ModuleWriter) WriteTypeAliases() {
	if len(m.FuncAliases) == 0 {
		return
	}

	filePath := filepath.Join(m.PlatformDir, m.Module.Package, "aliastypes.gen.go")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cw := &CodeWriter{Writer: f, IndentStr: "\t"}
	cw.WriteLine(AutoGeneratedMark)
	cw.WriteLine("package " + m.Module.Package)

	cw.WriteLine("import (")

	// imports
	var imports = set.New("unsafe")
	for _, fa := range m.FuncAliases {
		imports.AddSet(fa.GoImports())
	}
	cw.Indent()
	imports.ForEach(func(value string) {
		if value != "github.com/progrium/darwinkit/macos/objc" {
			cw.WriteLine("\"" + value + "\"")
		}
	})
	cw.UnIndent()
	cw.WriteLine(")")

	for _, fa := range m.FuncAliases {
		if fa.DocURL != "" {
			cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", fa.Description))
			cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", fa.DocURL))
		}
		cw.WriteLineF("type %s = %s", fa.GName, fa.Type.GoName(&m.Module, false))
	}
}

func (m *ModuleWriter) WriteStructs() {
	if len(m.Structs) == 0 {
		return
	}

	filePath := filepath.Join(m.PlatformDir, m.Module.Package, "structs.gen.go")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cw := &CodeWriter{Writer: f, IndentStr: "\t"}
	cw.WriteLine(AutoGeneratedMark)
	cw.WriteLine("package " + m.Module.Package)

	cw.WriteLine("import (")

	// imports
	var imports = set.New("unsafe")
	for _, fa := range m.FuncAliases {
		imports.AddSet(fa.GoImports())
	}
	cw.Indent()
	imports.ForEach(func(value string) {
		if value != "github.com/progrium/macdriver/macos/objc" {
			cw.WriteLine("\"" + value + "\"")
		}
	})
	cw.UnIndent()
	cw.WriteLine(")")

	for _, s := range m.Structs {
		// if Ref type, allias to unsafe.Pointer
		if strings.HasSuffix(s.Name, "Ref") {
			if s.DocURL != "" {
				cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", s.Description))
				cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", s.DocURL))
			}

			cw.WriteLineF("type %s unsafe.Pointer", s.GoName)
			continue
		}
	}
}

func shouldSkipFunction(f *Function) bool {
	if f.Deprecated {
		return true
	}
	if hasBlockParam(f.Parameters) {
		return true
	}
	if _, ok := map[string]bool{
		"CGDirectDisplayCopyCurrentMetalDevice":            true,
		"CGColorSpaceCreateWithPropertyList":               true,
		"CGDisplayIOServicePort":                           true,
		"CGGetEventTapList":                                true,
		"CGColorConversionInfoCreateFromListWithArguments": true,
	}[f.Name]; ok {
		return true
	}
	return false
}

// WriteFunctions writes the go code to call exposed functions.
func (m *ModuleWriter) WriteFunctions() {
	if len(m.Functions) == 0 {
		return
	}

	filePath := filepath.Join(m.PlatformDir, m.Module.Package, "functions.gen.go")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cw := &CodeWriter{Writer: f, IndentStr: "\t"}
	cw.WriteLine(AutoGeneratedMark)
	cw.WriteLine("package " + m.Module.Package)

	//TODO: determine imports from functions
	cw.WriteLineF(`// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import "%s"`, m.Module.Header)
	for _, f := range m.Functions {
		if shouldSkipFunction(f) {
			continue
		}
		f.WriteCSignature(&m.Module, cw)
	}
	cw.WriteLine(`import "C"`)
	cw.WriteLine("import (")

	// imports
	var imports = set.New("unsafe")
	for _, fa := range m.FuncAliases {
		imports.AddSet(fa.GoImports())
	}
	cw.Indent()
	imports.ForEach(func(value string) {
		if value == "github.com/progrium/macdriver/macos/"+m.Module.Package {
			return
		}
		// avoid cycles:
		if value != "github.com/progrium/macdriver/macos/objc" {
			cw.WriteLine("\"" + value + "\"")
		}
	})
	cw.UnIndent()
	cw.WriteLine(")")

	for _, f := range m.Functions {
		if shouldSkipFunction(f) {
			continue
		}
		f.WriteGoCallCode(&m.Module, cw)
	}
}

// WriteFunctionWrappers writes the objc code to wrap exposed functions.
// The cgo type system is unaware of objective c types so these wrappers must exist to allow
// us to call the functions and return appropritely.
func (m *ModuleWriter) WriteFunctionWrappers() {
	if len(m.Functions) == 0 {
		return
	}

	filePath := filepath.Join(m.PlatformDir, m.Module.Package, "functions.gen.m")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cw := &CodeWriter{Writer: f, IndentStr: "\t"}
	cw.WriteLine(AutoGeneratedMark)

	//TODO: determine appropriate imports
	cw.WriteLineF("#import \"%s\"", m.Module.Header)
	for _, f := range m.Functions {
		if shouldSkipFunction(f) {
			continue
		}
		f.WriteObjcWrapper(&m.Module, cw)
	}
}

func (m *ModuleWriter) WriteEnumAliases() {
	enums := make([]*AliasInfo, len(m.EnumAliases))
	copy(enums, m.EnumAliases)
	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	if len(enums) == 0 {
		return
	}
	armVals := false
	imports := set.New[string]()
	for _, e := range enums {
		imports.AddSet(e.Type.GoImports())
		for _, v := range e.Values {
			if v.Arm64Value != "" {
				armVals = true
			}
		}
	}

	//log.Println("gen enums for", m.Module.Name)

	filePath := filepath.Join(m.PlatformDir, m.Module.Package, "enumtypes.gen.go")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var arm64f io.Writer = io.Discard
	var amd64f io.Writer = io.Discard
	if armVals {
		armf, err := os.Create(filepath.Join(m.PlatformDir, m.Module.Package, "enumtypes_arm64.gen.go"))
		if err != nil {
			panic(err)
		}
		defer armf.Close()
		arm64f = armf
		amdf, err := os.Create(filepath.Join(m.PlatformDir, m.Module.Package, "enumtypes_amd64.gen.go"))
		if err != nil {
			panic(err)
		}
		defer amdf.Close()
		amd64f = amdf
	}

	cw := &CodeWriter{Writer: f, IndentStr: "\t"}
	cw.WriteLine(AutoGeneratedMark)
	cw.WriteLine("package " + m.Module.Package)

	arm64cw := &CodeWriter{Writer: arm64f, IndentStr: "\t"}
	arm64cw.WriteLine(AutoGeneratedMark)
	arm64cw.WriteLine("package " + m.Module.Package)

	amd64cw := &CodeWriter{Writer: amd64f, IndentStr: "\t"}
	amd64cw.WriteLine(AutoGeneratedMark)
	amd64cw.WriteLine("package " + m.Module.Package)

	cw.WriteLine("import (")
	cw.Indent()
	imports.ForEach(func(value string) {
		if value != "github.com/progrium/darwinkit/macos/objc" {
			cw.WriteLine("\"" + value + "\"")
		}
	})
	cw.UnIndent()
	cw.WriteLine(")")

	for _, ei := range enums {
		primitiveType := ei.Type.GoName(&m.Module, false)
		if ei.Module.Name == m.Module.Name {
			if ei.DocURL != "" {
				cw.WriteLine(fmt.Sprintf("// %s [Full Topic]", ei.Description))
				cw.WriteLine(fmt.Sprintf("//\n// [Full Topic]: %s", ei.DocURL))
			}
			if ei.Name == "CGFloat" {
				// special case CGFloat to be an actual type alias
				cw.WriteLine(fmt.Sprintf("type %s = %s\n", ei.GName, primitiveType))
			} else {
				cw.WriteLine(fmt.Sprintf("type %s %s\n", ei.GName, primitiveType))
			}
		}
		sort.Slice(ei.Values, func(i, j int) bool {
			return ei.Values[i].Name < ei.Values[j].Name
		})
		if len(ei.Values) > 0 {
			cw.WriteLine("const (")
			amd64cw.WriteLine("const (")
			arm64cw.WriteLine("const (")
		}
		var infValues []*EnumValue
		for _, v := range ei.Values {
			if v.Module == nil {
				continue
			}
			if v.Module.Name != m.Module.Name {
				// TODO: just skip it for now...
				continue
			}
			c := modules.LookupConstant("macos", m.Module.Package, v.Name)
			if c == nil {
				log.Println("MISSING CONSTANT:", v.Name)
			}
			if c != nil {
				v.Value = c.Value
				v.Arm64Value = c.ArmValue
			}
			if primitiveType == "uint" || primitiveType == "uint64" {
				// special cases maybe we wouldn't need if we exported properly
				if v.Value == "-1" {
					v.Value = "math.MaxUint"
					if v.Arm64Value != "" {
						v.Arm64Value = "math.MaxUint"
					}
				} else if strings.HasPrefix(v.Value, "-") {
					v.Value = strings.TrimPrefix(v.Value, "-")
					if v.Arm64Value != "" {
						v.Arm64Value = strings.TrimPrefix(v.Arm64Value, "-")
					}
				}
			}
			if v.Value == "inf" {
				infValues = append(infValues, v)
				continue
			}
			if v.Value == "" {
				log.Println("enum ", v.Name, " requires a value")
				continue
			}
			if v.GoName == ei.GoName(&m.Module, false) {
				v.GoName = "K" + v.GoName
			}
			if v.Arm64Value != "" {
				if !ei.IsString() {
					amd64cw.WriteLine(fmt.Sprintf("\t%s %s = %s", v.GoName, ei.GoName(&m.Module, false), v.Value))
					arm64cw.WriteLine(fmt.Sprintf("\t%s %s = %s", v.GoName, ei.GoName(&m.Module, false), v.Arm64Value))
				} else {
					amd64cw.WriteLine(fmt.Sprintf("\t%s %s = \"%s\"", v.GoName, ei.GoName(&m.Module, false), v.Value))
					arm64cw.WriteLine(fmt.Sprintf("\t%s %s = \"%s\"", v.GoName, ei.GoName(&m.Module, false), v.Arm64Value))
				}
			} else {
				if !ei.IsString() {
					cw.WriteLine(fmt.Sprintf("\t%s %s = %s", v.GoName, ei.GoName(&m.Module, false), v.Value))
				} else {
					cw.WriteLine(fmt.Sprintf("\t%s %s = \"%s\"", v.GoName, ei.GoName(&m.Module, false), v.Value))
				}
			}
		}
		if len(ei.Values) > 0 {
			cw.WriteLine(")")
			amd64cw.WriteLine(")")
			arm64cw.WriteLine(")")
		}
		for _, v := range infValues {
			cw.WriteLine(fmt.Sprintf("var %s = %s(math.Inf(1))", v.GoName, ei.GoName(&m.Module, false)))
		}
	}

}

func (m *ModuleWriter) WriteDocFile() {
	if m.DocURL != "" {
		f, err := os.Create(filepath.Join(m.PlatformDir, m.Module.Package, "doc.gen.go"))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		cw := &CodeWriter{Writer: f, IndentStr: "    "}
		cw.WriteLine(AutoGeneratedMark)
		cw.WriteLineF("// %s", m.Description)
		cw.WriteLine("//")
		cw.WriteLine("// [Apple Documentation]")
		cw.WriteLineF("//\n// [Apple Documentation]: %s", m.DocURL)
		cw.WriteLineF("package %s", m.Module.Package)
	}
}

func (m *ModuleWriter) WriteProtocolsImportCode() {
	//log.Println("gen protocol imports for", m.Module.Name)
	ps := make([]*typing.ProtocolType, len(m.Protocols))
	copy(ps, m.Protocols)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Name < ps[j].Name
	})
	m.writeProtocolMFile(ps, filepath.Join(m.PlatformDir, m.Module.Package, "protocols.gen.m"))
}

func (m *ModuleWriter) writeProtocolMFile(protocols []*typing.ProtocolType, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cw := &CodeWriter{Writer: f, IndentStr: "    "}
	cw.WriteLine(AutoGeneratedMark)
	cw.WriteLineF("#import \"%s\"", m.Module.Header)
	cw.WriteLine("")

	cw.WriteLineF("void import%sProtocols() {", m.Module.Name)
	cw.Indent()
	cw.WriteLine("id o;")
	for _, protocol := range protocols {
		cw.WriteLineF("o = @protocol(%s);", protocol.Name)
	}
	cw.UnIndent()
	cw.WriteLine("}")
}

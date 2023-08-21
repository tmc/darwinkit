package generate

import (
	"log"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
	"github.com/progrium/macdriver/generate/typing"
)

func (db *Generator) ToStruct(fw string, sym Symbol) *codegen.Struct {
	knownIssues := map[string]bool{}
	if knownIssues[sym.Name] {
		_, err := sym.Parse()
		log.Printf("skipping struct %s %s because of known issue: decl='%s' err='%v'\n", fw, sym.Name, sym.Declaration, err)
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	styp, ok := typ.(*typing.StructType)
	if !ok {
		return nil
	}
	s := &codegen.Struct{
		Name:        sym.Name,
		GoName:      modules.TrimPrefix(sym.Name),
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        styp,
	}

	return s

}

/*
 */
package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Kind describes the kind of the type (alias, array, etc.)
type Kind int

const (
	AliasKind Kind = iota
	ArrayKind
	BasicKind
	InterfaceKind
	MapKind
	PointerKind
	SliceKind
	StructKind
	UnknownKind
	UnsupportedKind
)

func (k *Kind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToUpper(s) {
	case "ALIAS":
		*k = AliasKind
	case "ARRAY":
		*k = ArrayKind
	case "BASIC":
		*k = BasicKind
	case "INTERFACE":
		*k = InterfaceKind
	case "MAP":
		*k = MapKind
	case "POINTER":
		*k = PointerKind
	case "SLICE":
		*k = SliceKind
	case "STRUCT":
		*k = StructKind
	case "UNKNOWN":
		*k = UnknownKind
	default:
		return fmt.Errorf("unknown kind %s", s)
	}
	return nil
}

func (k Kind) MarshalJSON() ([]byte, error) {
	kindStr := "UNKNOWN"
	switch k {
	case AliasKind:
		kindStr = "ALIAS"
	case ArrayKind:
		kindStr = "ARRAY"
	case BasicKind:
		kindStr = "BASIC"
	case InterfaceKind:
		kindStr = "INTERFACE"
	case MapKind:
		kindStr = "MAP"
	case PointerKind:
		kindStr = "POINTER"
	case SliceKind:
		kindStr = "SLICE"
	case StructKind:
		kindStr = "STRUCT"
	}

	return json.Marshal(kindStr)
}

// Type describes a declared type
type Type struct {
	Name           string                   `json:"name"`
	Package        string                   `json:"package"`
	Doc            string                   `json:"doc"`
	GVK            *schema.GroupVersionKind `json:"gvk"`
	Kind           Kind                     `json:"kind"`
	Imported       bool                     `json:"imported"`
	UnderlyingType *Type                    `json:"underlyingType"` // for aliases, slices and pointers
	KeyType        *Type                    `json:"keyType"`        // for maps
	ValueType      *Type                    `json:"valueType"`      // for maps
	Fields         []*Field                 `json:"fields"`         // for structs
	References     []*Type                  `json:"-"`              // other types that refer to this type
}

func (t *Type) IsBasic() bool {
	switch t.Kind {
	case BasicKind:
		return true
	case AliasKind, SliceKind, ArrayKind, PointerKind:
		return t.UnderlyingType != nil && t.UnderlyingType.IsBasic()
	case MapKind:
		return t.KeyType != nil && t.KeyType.IsBasic() && t.ValueType != nil && t.ValueType.IsBasic()
	case InterfaceKind:
		return true
	default:
		return false
	}
}

func (t *Type) Members() []*Field {
	if t == nil {
		return nil
	}

	if len(t.Fields) > 0 {
		return t.Fields
	}

	switch t.Kind {
	case AliasKind, SliceKind, ArrayKind, PointerKind:
		return t.UnderlyingType.Members()
	default:
		return nil
	}
}

func (t *Type) String() string {
	if t == nil {
		return "<unknown>"
	}

	var sb strings.Builder
	switch t.Kind {
	case MapKind:
		sb.WriteString("map[")
		sb.WriteString(t.KeyType.String())
		sb.WriteString("]")
		sb.WriteString(t.ValueType.String())
		return sb.String()
	case ArrayKind, SliceKind:
		sb.WriteString("[]")
	case PointerKind:
		sb.WriteString("*")
	}
	if t.Package != "" {
		sb.WriteString(t.Package)
		sb.WriteString(".")
	}
	sb.WriteString(t.Name)

	return sb.String()
}

func (t *Type) IsAlias() bool {
	return t.Kind == AliasKind
}

// Field describes a field in a struct.
type Field struct {
	Name     string
	Embedded bool
	Doc      string
	Type     *Type
}

// Key generates the unique name for the give type.
func Key(t *Type) string {
	if t.Package == "" {
		return t.Name
	}
	return fmt.Sprintf("%s.%s", t.Package, t.Name)
}

// GroupVersionDetails encapsulates details about a discovered API group.
type GroupVersionDetails struct {
	schema.GroupVersion
	Kinds []string
	Types map[string]*Type
}

func (gvd GroupVersionDetails) GroupVersionString() string {
	return gvd.GroupVersion.String()
}

func (gvd GroupVersionDetails) TypeForKind(k string) *Type {
	return gvd.Types[k]
}

func (gvd GroupVersionDetails) SortedTypes() []*Type {
	typeList := make([]*Type, len(gvd.Types))
	i := 0
	for _, t := range gvd.Types {
		typeList[i] = t
		i++
	}

	sort.Slice(typeList, func(i, j int) bool {
		return typeList[i].Name < typeList[j].Name
	})

	return typeList
}

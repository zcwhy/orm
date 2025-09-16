package dialect

import "reflect"

var dialectsMap = map[string]Dialector{}

type Dialector interface {
	DataTypeOf(typ reflect.Type) string
}

func init() {
	RegisterDialector("sqlite3", &SqlliteDialect{})
}

func RegisterDialector(name string, dialect Dialector) {
	dialectsMap[name] = dialect
}

func GetDialector(name string) (dialect Dialector, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}

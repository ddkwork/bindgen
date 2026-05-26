package cpp2c

import (
	"testing"
)

func TestProcessFreeFunction_BlocksQtInternalFunctions(t *testing.T) {
	tests := []struct {
		name        string
		methodName  string
		qualType    string
		shouldBlock bool
	}{
		{"qt_noop should be blocked", "qt_noop", "void ()", true},
		{"qt_assert should be blocked", "qt_assert", "void (const char *, const char *, int)", true},
		{"qt_assert_x should be blocked", "qt_assert_x", "void (const char *, const char *, const char *, int)", true},
		{"qt_no_assert_x should be blocked", "qt_no_assert_x", "bool (bool, const char *, const char *)", true},
		{"qt_check_pointer should be blocked", "qt_check_pointer", "void (const char *, int)", true},
		{"qBadAlloc should be blocked", "qBadAlloc", "void ()", true},
		{"qt_assume_is_deprecated should be blocked", "qt_assume_is_deprecated", "bool (bool)", true},
		{"qInstallMessageHandler should be blocked", "qInstallMessageHandler", "QtMessageHandler (QtMessageHandler)", true},
		{"qt_error_string should be blocked", "qt_error_string", "miqt_string (int)", true},
		{"qgetenv should be blocked", "qgetenv", "miqt_string (const char *)", true},
		{"qunsetenv should be blocked", "qunsetenv", "bool (const char *)", true},
		{"qtTrId should be blocked", "qtTrId", "miqt_string (const char *, int)", true},
		{"qEnvironmentVariableIsEmpty should be blocked", "qEnvironmentVariableIsEmpty", "bool (const char *)", true},
		{"qEnvironmentVariableIsSet should be blocked", "qEnvironmentVariableIsSet", "bool (const char *)", true},
		{"qEnvironmentVariableIntValue should be blocked", "qEnvironmentVariableIntValue", "int (const char *, bool *)", true},
		{"qEnvironmentVariableIntegerValue should be blocked", "qEnvironmentVariableIntegerValue", "void * (const char *)", true},
		{"qSNaN should be blocked", "qSNaN", "float ()", true},
		{"qQNaN should be blocked", "qQNaN", "float ()", true},
		{"qInf should be blocked", "qInf", "float ()", true},
		{"qMallocAligned should be blocked", "qMallocAligned", "void * (size_t, size_t)", true},
		{"qReallocAligned should be blocked", "qReallocAligned", "void * (void *, size_t, size_t, size_t)", true},
		{"qFreeAligned should be blocked", "qFreeAligned", "void (void *)", true},
		{"qt_getEnumMetaObject_1 should be blocked", "qt_getEnumMetaObject_1", "QMetaObject * (ScrollBarPolicy)", true},
		{"qt_getEnumName_1 should be blocked", "qt_getEnumName_1", "const char * (ScrollBarPolicy)", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &AstNode{
				Kind: "FunctionDecl",
				Fields: map[string]interface{}{
					"name": tt.methodName,
					"type": map[string]interface{}{
						"qualType": tt.qualType,
					},
				},
				Inner: nil,
			}

			result, err := processFreeFunction(node, "")

			if tt.shouldBlock {
				if err == nil {
					t.Errorf("processFreeFunction(%q) expected to be blocked but got result: %+v", tt.methodName, result)
				}
			} else {
				if err != nil {
					t.Errorf("processFreeFunction(%q) should NOT be blocked but got error: %v", tt.methodName, err)
				}
			}
		})
	}
}

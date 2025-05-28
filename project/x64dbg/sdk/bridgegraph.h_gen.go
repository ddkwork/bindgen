package sdk

// BridgeCFNodeList (:10)
type BridgeCFNodeList struct {
	ParentGraph  int      // C type: int
	Start        int      // C type: int
	End          int      // C type: int
	Brtrue       int      // C type: int
	Brfalse      int      // C type: int
	Icount       int      // C type: int
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        int      // C type: int
	Instrs       int      // C type: int
}

// BridgeCFNode (D:\workspace\workspace\mcp\x64dbg\include\bridgegraph.h:41)
type BridgeCFNode struct {
	ParentGraph  int      // C type: int
	Start        int      // C type: int
	End          int      // C type: int
	Brtrue       int      // C type: int
	Brfalse      int      // C type: int
	Icount       int      // C type: int
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        int      // C type: int
	Instrs       uintptr  // C type: std::vector<BridgeCFInstruction>
}

// BridgeCFGraph (:112)
type BridgeCFGraph struct {
	EntryPoint int      // C type: int
	Userdata   *uintptr // C type: void *
	Nodes      int      // C type: int
}
type bridgegraph struct{}

func (b *bridgegraph) Free(graphList *uintptr) {
	response := safeGet("bridgegraph.h/Free", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgegraph) ToNodeList() {
	response := safeGet("bridgegraph.h/ToNodeList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgegraph) __debugbreak() {
	response := safeGet("bridgegraph.h/__debugbreak", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgegraph) AddNode(node *BridgeCFNode) {
	response := safeGet("bridgegraph.h/AddNode", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgegraph) AddParent(child int, parent int) {
	response := safeGet("bridgegraph.h/AddParent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgegraph) ToGraphList() {
	response := safeGet("bridgegraph.h/ToGraphList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

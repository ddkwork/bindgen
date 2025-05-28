package sdk

type stack struct{}

func (s *stack) Pop() {
	response := safeGet("_scriptapi_stack.h/Pop", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (s *stack) Push(value uint) {
	response := safeGet("_scriptapi_stack.h/Push", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (s *stack) Peek(offset int) {
	response := safeGet("_scriptapi_stack.h/Peek", map[string]string{})
	if len(response) == 0 {
		return
	}
}

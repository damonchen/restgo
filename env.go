package restgo

// Environment for test stage running
type Environment struct {
	context map[string]interface{}
}

func (e *Environment) Has(key string) bool {
	if _, ok := e.context[key]; ok {
		return true
	}
	return false
}

func (e *Environment) Get(key string) interface{} {
	return e.context[key]
}

func (e *Environment) Set(key string, value interface{}) {
	e.context[key] = value
}

func NewEnvironment() *Environment {
	return &Environment{
		context: map[string]interface{}{},
	}
}

package prototype

import "fmt"

type Registry struct {
	RegisterMap map[string]Shape
}

func NewRegistry() *Registry {
	return &Registry{
		RegisterMap: make(map[string]Shape),
	}
}

func (r *Registry) Register(name string, shape Shape) {
	r.RegisterMap[name] = shape
}

func (r *Registry) Get(name string) (Shape, error) {
	prototype, exists := r.RegisterMap[name]
	if !exists {
		return nil, fmt.Errorf("no such shape: %s", name)
	}

	return prototype.Clone(), nil

}

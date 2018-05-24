package main

type Resource struct {
	name string
}

type Resources struct {
	res []*Resource
}

func NewResource(name string) *Resource {
	return &Resource{name: name}
}

func (r *Resource) Name() string {
	return r.name
}

func NewResources() *Resources {
	return &Resources{}
}

func (r *Resources) ToStringSlice() []string {
	var list []string
	for _, res := range r.res {
		list = append(list, res.Name())
	}
	return list
}

func (r *Resources) Add(res *Resource) {
	r.res = append(r.res, res)
}

func (r *Resources) NewResource(name string) *Resource {
	res := NewResource(name)
	r.Add(res)
	return res
}

func (r *Resources) Has(res *Resource) bool {
	for _, resource := range r.res {
		if resource.Name() == res.Name() {
			return true
		}
	}
	return false
}

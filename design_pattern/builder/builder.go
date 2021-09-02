package builder

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type Builder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func newResourcePoolConfig(b *Builder) *ResourcePoolConfig {
	r := &ResourcePoolConfig{}

	r.name = b.name
	r.maxTotal = b.maxTotal
	r.maxIdle = b.maxIdle
	r.minIdle = b.minIdle
	return r
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() *ResourcePoolConfig {
	if b.maxIdle > b.maxTotal {
		panic("maxIdle must be less than maxTotal")
	}

	if b.minIdle > b.maxIdle || b.minIdle > b.maxTotal {
		panic("minIdle must be less than maxIdle or maxTotal")
	}

	return newResourcePoolConfig(b)
}

func (b *Builder) SetName(name string) *Builder {
	if len(name) == 0 {
		panic("name should not be empty")
	}
	b.name = name
	return b
}

func (b *Builder) SetMaxTotal(maxTotal int) *Builder {
	if maxTotal <= 0 {
		panic("maxTotal should be positive")
	}

	b.maxTotal = maxTotal
	return b
}
func (b *Builder) SetMaxIdle(maxIdle int) *Builder {
	if maxIdle <= 0 {
		panic("maxIdle should be positive")
	}
	b.maxIdle = maxIdle
	return b
}

func (b *Builder) SetMinIdle(minIdle int) *Builder {
	if minIdle <= 0 {
		panic("minIdle should be positive")
	}
	b.minIdle = minIdle
	return b
}

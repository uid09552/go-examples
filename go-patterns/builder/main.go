package builder

type Builder struct {
	entity   Entity
	property string
	name     string
}
type Entity struct {
	Name     string
	Property string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b Builder) Build() *Entity {
	return &Entity{Name: b.name, Property: b.property}
}

func (b Builder) Property1(value string) *Builder {
	b.entity.Property = value
	return &b
}

func (b Builder) Name(value string) *Builder {
	b.entity.Name = value
	return &b
}

func main() {
	builder := NewBuilder()
	entity := builder.Name("test").Property1("test").Build()
	print(entity)

}

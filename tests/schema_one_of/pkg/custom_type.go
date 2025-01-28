package pkg

type Dog struct {
	Name string
	Tag  Maybe[string]
}

type DogSchema struct {
	Name string
	Tag  Maybe[string]
}

func (d *Dog) SetFromSchemaDog(s DogSchema) error {
	*d = Dog{
		Name: s.Name,
		Tag:  s.Tag,
	}
	return nil
}

func (d Dog) ToSchemaDog() DogSchema {
	return DogSchema{
		Name: d.Name,
		Tag:  d.Tag,
	}
}

type Dog2 struct {
	Name    Maybe[string]
	PetType string
	Tag     Maybe[string]
}

type DogSchema2 struct {
	Name    Maybe[string]
	PetType string
	Tag     Maybe[string]
}

func (d *Dog2) SetFromSchemaDog2(s DogSchema2) error {
	*d = Dog2(s)
	return nil
}

func (d Dog2) ToSchemaDog2() DogSchema2 {
	return DogSchema2(d)
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

func Pointer[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Nullable[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

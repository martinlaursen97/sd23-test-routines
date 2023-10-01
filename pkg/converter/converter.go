package converter

type Converter interface {
	Convert(...any) any
}

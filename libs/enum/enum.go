package enum

// EnumValid defines an interface for validating enum types.
type EnumValid interface {
	Valid() bool
}

// EnumStringable defines an interface for converting enums to and from string representations.
type EnumStringable interface {
	FromString(str string) error
}

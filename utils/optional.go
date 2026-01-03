package utils

func ApplyOptionalValue[T any](isSet func() bool, getter func() T, target *T) {
	if isSet() {
		*target = getter()
	}
}

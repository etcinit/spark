package core

// Truthy returns the first true result from a slice of functions that return
// bools. Each function is only called when the element is being evaluated.
// This means that some function might not be called if an earlier element
// returns true.
//
// The default value is **false**.
func Truthy(xs []func() bool) bool {
	for _, value := range xs {
		if value() == true {
			return true
		}
	}

	return false
}

// Falsy returns the first false result from a slice of functions that return
// bools. Each function is only called when the element is being evaluated.
// This means that some function might not be called if an earlier element
// returns false.
//
// The default value is **false**.
func Falsy(xs []func() bool) bool {
	for _, value := range xs {
		if value() == false {
			return false
		}
	}

	return true
}

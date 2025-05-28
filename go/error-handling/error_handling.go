package erratum

// Note: The types ResourceOpener, Resource, TransientError, and FrobError
// are defined in common.go and are available in this package.

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	// Loop to open the resource, retrying on TransientError.
	for {
		res, openErr := opener()
		if openErr != nil {
			if _, ok := openErr.(TransientError); ok {
				// It's a TransientError, so retry opening.
				continue
			}
			// It's a non-transient error; return it immediately.
			return openErr
		}
		resource = res // Successfully opened the resource.
		break         // Exit the opener loop.
	}

	// Defer Close to ensure it's always called if the resource was successfully opened.
	// This defer will execute after the panic recovery defer (if a panic occurs).
	defer func() {
		// At this point, 'resource' is guaranteed to be non-nil because we broke
		// out of the loop above only upon successful opening.
		resource.Close()
	}()

	// Defer panic handling for Frob. This defer executes first (LIFO) if Frob panics.
	defer func() {
		if r := recover(); r != nil {
			// A panic occurred during resource.Frob().
			if frobErr, ok := r.(FrobError); ok {
				// It's a FrobError. Call Defrob and set the error to be returned.
				resource.Defrob(frobErr.defrobTag)
				err = frobErr // Set the named return error for Use().
			} else if e, ok := r.(error); ok {
				// It's another error type that was panicked.
				err = e // Set the named return error.
			} else {
				// It's some other unknown panic type. Re-panic it, as Use()
				// is defined to return an error, not handle arbitrary panic values directly.
				panic(r)
			}
		}
	}()

	// Call Frob. This might panic.
	resource.Frob(input)

	// If Frob completed without panic, 'err' (the named return) is still nil (its zero value).
	return err
}

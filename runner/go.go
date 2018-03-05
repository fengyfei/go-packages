package runner

// Go starts a goroutine with the default recover handler.
func Go(goroutine func()) {
	GoWithRecover(goroutine, nil)
}

// GoWithRecover starts a goroutine with a custom recover handler.
// If recoverHandler is nil, a default handler is used to print the call stack.
func GoWithRecover(goroutine func(), recoverHandler func(interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if recoverHandler != nil {
					recoverHandler(err)
				} else {
					defaultRecoverHandler(err)
				}
			}
		}()

		goroutine()
	}()
}

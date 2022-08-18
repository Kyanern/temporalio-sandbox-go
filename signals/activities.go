package signals

import (
	"context"
	"log"
)

func Activity_stub1(_ context.Context) (string, error) {
	const FUNCNAME = "Activity_stub1"
	log.Println("Entered " + FUNCNAME + "...")
	log.Println("Leaving " + FUNCNAME + ". Returning a string.")

	var retval string = FUNCNAME + " returned this string."
	return retval, nil
}

func Activity_stub2(_ context.Context) (string, error) {
	const FUNCNAME = "Activity_stub2"
	log.Println("Entered " + FUNCNAME + "...")
	log.Println("Leaving " + FUNCNAME + ". Returning a string.")

	var retval string = FUNCNAME + " returned this string."
	return retval, nil
}

// func Activity_waitForSignal(ctx context.Context)(string, error){
// 	const FUNCNAME = "waitForSignal"
// 	log.Println("Entered " + FUNCNAME + "...")

// 	//code to wait for signal here maybe? nope. activity doesn't have NewSelector.

// 	log.Println("Leaving " + FUNCNAME + ". Returning a string.")

// 	var retval string = FUNCNAME + " returned this string."
// 	return retval, nil
// }

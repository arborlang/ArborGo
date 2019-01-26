package environment

import (
	"fmt"

	"github.com/arborlang/arbor-dev"
)

// RunWasm runs a Wasm file
func RunWasm(wasmCode []byte, entrypoint string, paths ...string) (int64, error) {
	vm, err := arbor.NewVirtualMachine(wasmCode, entrypoint, paths...)
	if err != nil {
		return int64(-1), err
	}
	if err != nil {
		return int64(-1), err
	}
	ret, err := vm.Run()
	if err != nil {
		vm.PrintStackTrace()
		return int64(-1), err
	}
	return ret, nil
}

// RunWat runs a Wat file
func RunWat() (int64, error) {
	return int64(-1), fmt.Errorf("not implemented")
}

//RunArbor runs an arbor file
func RunArbor(wasmCode []byte, entrypoint string) (int64, error) {
	return int64(-1), fmt.Errorf("not implemented")
}

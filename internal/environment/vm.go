package environment

import (
	"fmt"
	"github.com/perlin-network/life/exec"
)

// RunWasm runs a Wasm file
func RunWasm(wasmCode []byte, entrypoint string) (int64, error) {
	vm, err := exec.NewVirtualMachine(wasmCode, exec.VMConfig{}, &exec.NopResolver{}, nil)
	if err != nil {
		return int64(-1), err
	}
	entryID, ok := vm.GetFunctionExport(entrypoint) // can be changed to your own exported function
	if !ok {
		return int64(-1), fmt.Errorf("entry function not found")
	}
	ret, err := vm.Run(entryID)
	if err != nil {
		vm.PrintStackTrace()
		return int64(-1), err
	}
	return ret, nil
	// return int64(-1), fmt.Errorf("not implemented")
}

// RunWat runs a Wat file
func RunWat() (int64, error) {
	return int64(-1), fmt.Errorf("not implemented")
}

//RunArbor runs an arbor file
func RunArbor(wasmCode []byte, entrypoint string) (int64, error) {
	return int64(-1), fmt.Errorf("not implemented")
}

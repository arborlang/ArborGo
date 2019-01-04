package environment

import (
	"github.com/perlin-network/life/exec"
)

// RunWasm runs a Wasm file
func RunWasm(wasmCode []byte, entrypoint string) (int, error) {
	vm, err := exec.NewVirtualMachine(wasmCode, exec.VMConfig{}, &exec.NopResolver{})
	if err != nil {
		return -1, err
	}
	entryID, ok := vm.GetFunctionExport(entrypoint) // can be changed to your own exported function
	if !ok {
		return -1, fmt.Errorf("entry function not found")
	}
	ret, err := vm.Run(entryID)
	if err != nil {
		vm.PrintStackTrace()
		return -1, err
	}
	return ret, nil
}

// RunWat runs a Wat file
func RunWat() error {
	return fmt.Errorf("not implemented")
}

//RunArbor runs an arbor file
func RunArbor(wasmCode []byte, entrypoint string) (int, error) {
	return fmt.Errorf("not implemented")
}

package environment

import (
	"fmt"
	"github.com/perlin-network/life/exec"
)

//ResolverManager handles all of the resolvers
type ResolverManager struct {
	resolvers map[string]func(vm *exec.VirtualMachine) int64
}

type SysResolver struct{}

func (r *SysResolver) ResolveFunc(module, field string) exec.FunctionImport {
	switch module {
	case "env":
		switch field {
		case "__putch__":
			return func(vm *exec.VirtualMachine) int64 {
				ptr := rune(vm.GetCurrentFrame().Locals[0])
				fmt.Print(string(ptr))
				return 0
			}

		default:
			panic(fmt.Errorf("unknown import resolved: %s", field))
		}
	default:
		panic(fmt.Errorf("unknown module: %s", module))
	}
}

func (r *SysResolver) ResolveGlobal(module, field string) int64 {
	panic("we're not resolving global variables for now")
}

package environment

//ResolverManager handles all of the resolvers
type ResolverManager struct {
	resolvers map[string]func(vm *exec.VirtualMachine) int64
}

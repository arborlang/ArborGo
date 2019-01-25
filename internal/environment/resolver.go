package environment

import (
	"fmt"
	"github.com/perlin-network/life/exec"
	"github.com/radding/arbor-dev"
	"plugin"
)

//ResolverManager handles all of the resolvers
type ResolverManager struct {
	resolvers map[string]arbor.Module
}

//NewResolver makes a resolver
func NewResolver() *ResolverManager {
	manager := new(ResolverManager)
	manager.resolvers = map[string]arbor.Module{}
	return manager
}

// Load loads a module
func (r *ResolverManager) Load(name string) error {
	plug, err := plugin.Open(name)
	if err != nil {
		return err
	}
	resolver, err := plug.Lookup("Env")
	if err != nil {
		return err
	}
	if module, ok := resolver.(arbor.Module); ok {
		r.resolvers[module.Name()] = module
		return nil
	}
	return fmt.Errorf("Could not open extentions")
}

//ResolveFunc finds the function you are looking for
func (r *ResolverManager) ResolveFunc(module, field string) exec.FunctionImport {
	mod, ok := r.resolvers[module]
	if !ok {
		panic(fmt.Errorf("unknown import resolved: %s", module))
	}
	ext := mod.Resolve(field)
	if ext == nil {
		panic(fmt.Errorf("%s has no function %s", module, field))
	}
	return ext.Run
}

//ResolveGlobal just dies
func (r *ResolverManager) ResolveGlobal(module, field string) int64 {
	panic("we're not resolving global variables for now")
}

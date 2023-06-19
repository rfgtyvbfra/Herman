package execution

// EVMEngine v0 provides a mock execution engine.
type EVMEngineV0 struct{}

func (EVMEngineV0) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV0) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm0-"), tx...)
    return out, nil
}

import "fmt"
package execution

// EVMEngine v4 provides a mock execution engine.
type EVMEngineV4 struct{}

func (EVMEngineV4) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV4) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm4-"), tx...)
    return out, nil
}

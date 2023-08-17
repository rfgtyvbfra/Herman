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
package execution

// EVMEngine v8 provides a mock execution engine.
type EVMEngineV8 struct{}

func (EVMEngineV8) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV8) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm8-"), tx...)
    return out, nil
}
package execution

// EVMEngine v12 provides a mock execution engine.
type EVMEngineV12 struct{}

func (EVMEngineV12) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV12) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm12-"), tx...)
    return out, nil
}
package execution

// EVMEngine v16 provides a mock execution engine.
type EVMEngineV16 struct{}

func (EVMEngineV16) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV16) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm16-"), tx...)
    return out, nil
}
package execution

// EVMEngine v20 provides a mock execution engine.
type EVMEngineV20 struct{}

func (EVMEngineV20) Init(vm VMType) error {
    if vm != VM_EVM {
        return fmt.Errorf("unsupported VM: %s", vm)
    }
    return nil
}

func (EVMEngineV20) Execute(tx []byte) ([]byte, error) {
    // Mock: echo back with prefix
    out := append([]byte("evm20-"), tx...)
    return out, nil
}

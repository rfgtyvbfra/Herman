package execution

type VMType string

const (
    VM_EVM VMType = "EVM"
    VM_WASM VMType = "WASM"
)

type ExecutionEngine interface {
    Init(vm VMType) error
    Execute(tx []byte) ([]byte, error)
}

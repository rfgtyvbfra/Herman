package settlement

// ProverV1 is a mock prover for settlement proofs.
type ProverV1 struct{}

func (ProverV1) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}

import "fmt"

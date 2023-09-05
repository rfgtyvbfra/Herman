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
package settlement

// ProverV5 is a mock prover for settlement proofs.
type ProverV5 struct{}

func (ProverV5) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}
package settlement

// ProverV9 is a mock prover for settlement proofs.
type ProverV9 struct{}

func (ProverV9) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}
package settlement

// ProverV13 is a mock prover for settlement proofs.
type ProverV13 struct{}

func (ProverV13) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}
package settlement

// ProverV17 is a mock prover for settlement proofs.
type ProverV17 struct{}

func (ProverV17) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}
package settlement

// ProverV21 is a mock prover for settlement proofs.
type ProverV21 struct{}

func (ProverV21) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}
package settlement

// ProverV25 is a mock prover for settlement proofs.
type ProverV25 struct{}

func (ProverV25) Finalize(stateRoot string) error {
    if len(stateRoot) == 0 {
        return fmt.Errorf("empty state root")
    }
    return nil
}

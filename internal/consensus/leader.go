package consensus

import "math/rand"

// LeaderElectorV2 chooses a pseudo-random leader.
type LeaderElectorV2 struct{}

func (LeaderElectorV2) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV2) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}

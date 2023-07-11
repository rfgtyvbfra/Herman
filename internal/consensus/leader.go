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
package consensus

import "math/rand"

// LeaderElectorV6 chooses a pseudo-random leader.
type LeaderElectorV6 struct{}

func (LeaderElectorV6) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV6) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV10 chooses a pseudo-random leader.
type LeaderElectorV10 struct{}

func (LeaderElectorV10) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV10) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}

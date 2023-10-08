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
package consensus

import "math/rand"

// LeaderElectorV14 chooses a pseudo-random leader.
type LeaderElectorV14 struct{}

func (LeaderElectorV14) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV14) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV18 chooses a pseudo-random leader.
type LeaderElectorV18 struct{}

func (LeaderElectorV18) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV18) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV22 chooses a pseudo-random leader.
type LeaderElectorV22 struct{}

func (LeaderElectorV22) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV22) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV26 chooses a pseudo-random leader.
type LeaderElectorV26 struct{}

func (LeaderElectorV26) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV26) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV30 chooses a pseudo-random leader.
type LeaderElectorV30 struct{}

func (LeaderElectorV30) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV30) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}
package consensus

import "math/rand"

// LeaderElectorV34 chooses a pseudo-random leader.
type LeaderElectorV34 struct{}

func (LeaderElectorV34) Start(validators []string) error {
    if len(validators) == 0 {
        return fmt.Errorf("no validators")
    }
    return nil
}

func (LeaderElectorV34) Propose(block []byte) error {
    _ = rand.Int() // simulate work
    if len(block) == 0 {
        return fmt.Errorf("empty block")
    }
    return nil
}

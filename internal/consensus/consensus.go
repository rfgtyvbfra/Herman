package consensus

type Algorithm string

const (
    PoS Algorithm = "PoS"
    PoA Algorithm = "PoA"
)

type Consensus interface {
    Start(validators []string) error
    Propose(block []byte) error
}

package settlement

type SettlementLayer interface {
    Finalize(stateRoot string) error
}

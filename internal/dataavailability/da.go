package dataavailability

type DataAvailability interface {
    Store(chunk []byte) (string, error)
    Sample(cid string, n int) ([][]byte, error)
}

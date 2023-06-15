package dataavailability

// MemStoreV3 stores chunks in-memory for tests.
type MemStoreV3 struct{
    data map[string][]byte
}

func (m *MemStoreV3) Store(chunk []byte) (string, error) {
    if m.data == nil { m.data = make(map[string][]byte) }
    key := fmt.Sprintf("cid-%d", len(m.data)+3)
    m.data[key] = chunk
    return key, nil
}

func (m *MemStoreV3) Sample(cid string, n int) ([][]byte, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0")
    }
    out := make([][]byte, n)
    for i := 0; i < n; i++ {
        out[i] = []byte(cid)
    }
    return out, nil
}

import "fmt"

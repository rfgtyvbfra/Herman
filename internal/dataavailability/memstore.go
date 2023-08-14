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
package dataavailability

// MemStoreV7 stores chunks in-memory for tests.
type MemStoreV7 struct{
    data map[string][]byte
}

func (m *MemStoreV7) Store(chunk []byte) (string, error) {
    if m.data == nil { m.data = make(map[string][]byte) }
    key := fmt.Sprintf("cid-%d", len(m.data)+7)
    m.data[key] = chunk
    return key, nil
}

func (m *MemStoreV7) Sample(cid string, n int) ([][]byte, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0")
    }
    out := make([][]byte, n)
    for i := 0; i < n; i++ {
        out[i] = []byte(cid)
    }
    return out, nil
}
package dataavailability

// MemStoreV11 stores chunks in-memory for tests.
type MemStoreV11 struct{
    data map[string][]byte
}

func (m *MemStoreV11) Store(chunk []byte) (string, error) {
    if m.data == nil { m.data = make(map[string][]byte) }
    key := fmt.Sprintf("cid-%d", len(m.data)+11)
    m.data[key] = chunk
    return key, nil
}

func (m *MemStoreV11) Sample(cid string, n int) ([][]byte, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0")
    }
    out := make([][]byte, n)
    for i := 0; i < n; i++ {
        out[i] = []byte(cid)
    }
    return out, nil
}
package dataavailability

// MemStoreV15 stores chunks in-memory for tests.
type MemStoreV15 struct{
    data map[string][]byte
}

func (m *MemStoreV15) Store(chunk []byte) (string, error) {
    if m.data == nil { m.data = make(map[string][]byte) }
    key := fmt.Sprintf("cid-%d", len(m.data)+15)
    m.data[key] = chunk
    return key, nil
}

func (m *MemStoreV15) Sample(cid string, n int) ([][]byte, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0")
    }
    out := make([][]byte, n)
    for i := 0; i < n; i++ {
        out[i] = []byte(cid)
    }
    return out, nil
}
package dataavailability

// MemStoreV19 stores chunks in-memory for tests.
type MemStoreV19 struct{
    data map[string][]byte
}

func (m *MemStoreV19) Store(chunk []byte) (string, error) {
    if m.data == nil { m.data = make(map[string][]byte) }
    key := fmt.Sprintf("cid-%d", len(m.data)+19)
    m.data[key] = chunk
    return key, nil
}

func (m *MemStoreV19) Sample(cid string, n int) ([][]byte, error) {
    if n <= 0 {
        return nil, fmt.Errorf("n must be > 0")
    }
    out := make([][]byte, n)
    for i := 0; i < n; i++ {
        out[i] = []byte(cid)
    }
    return out, nil
}

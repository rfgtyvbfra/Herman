package execution

func checksumV1(b []byte) int {
    s := 0
    for _, v := range b { s += int(v) }
    return s
}
package execution

func checksumV3(b []byte) int {
    s := 0
    for _, v := range b { s += int(v) }
    return s
}

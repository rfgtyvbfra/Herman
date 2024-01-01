package execution

import "testing"

func TestExecuteEchoV16(t *testing.T) {
    eng := EVMEngineV16{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

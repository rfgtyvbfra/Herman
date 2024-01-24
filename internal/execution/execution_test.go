package execution

import "testing"

func TestExecuteEchoV22(t *testing.T) {
    eng := EVMEngineV22{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

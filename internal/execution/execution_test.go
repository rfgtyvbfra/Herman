package execution

import "testing"

func TestExecuteEchoV26(t *testing.T) {
    eng := EVMEngineV26{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

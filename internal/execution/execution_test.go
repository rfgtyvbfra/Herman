package execution

import "testing"

func TestExecuteEchoV12(t *testing.T) {
    eng := EVMEngineV12{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

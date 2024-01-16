package execution

import "testing"

func TestExecuteEchoV20(t *testing.T) {
    eng := EVMEngineV20{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

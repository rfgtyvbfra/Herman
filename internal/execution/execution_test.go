package execution

import "testing"

func TestExecuteEchoV10(t *testing.T) {
    eng := EVMEngineV10{}
    _ = eng.Init(VM_EVM)
    out, _ := eng.Execute([]byte("tx"))
    if len(out) == 0 { t.Fatal("no output") }
}

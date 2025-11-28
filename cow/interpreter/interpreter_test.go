package interpreter

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestParseCode(t *testing.T) {
	i := Interpreter()
	i.parseCode("MOO mOo comment moo")
	expectedCode := make([]command, 3)
	expectedCode[0] = command_MOO
	expectedCode[1] = command_mOo
	expectedCode[2] = command_moo
	for pos, expectedCommand := range expectedCode {
		if i.code[pos] != expectedCommand {
			t.Errorf("Got unexpected command %d instead of %d", i.code[pos], expectedCommand)
		}
	}
}

func TestExecuteCommand(t *testing.T) {
	i := Interpreter()
	oldPtr := i.ptr
	i.executeCommand(command_moO)
	if i.ptr != oldPtr+1 {
		t.Errorf("Incorrect command execution")
	}
}
func TestInterpret(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	filepath := filepath.Join("..", "cow_examples", "hello.cow")
	source, _ := os.ReadFile(filepath)
	i := Interpreter()
	err := i.Interpret(string(source))

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	expectedOutput := "Hello, World!"
	if err != nil {
		t.Errorf("Interpreter interrupted with error: %s", err)
	}
	if output != expectedOutput {
		t.Errorf("Got unexpected output %s instead of %s", output, expectedOutput)
	}
}

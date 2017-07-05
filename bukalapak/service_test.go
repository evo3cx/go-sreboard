package bukalapak

import "testing"

func TestRunCommand(t *testing.T) {
	_, err := runCommand("ls")
	if err != nil {
		t.Fatalf("expect nil get err %s", err)
	}

}

func TestSplitWithCommand(t *testing.T) {
	testCase := []struct {
		in      string
		outCmd  string
		outArgs []string
	}{
		{
			in:      "cat_service.go",
			outCmd:  "cat",
			outArgs: []string{"service.go"},
		},

		{
			in:      "ls_-la",
			outCmd:  "ls",
			outArgs: []string{"-la"},
		},
	}

	for i, tc := range testCase {
		cmd, args := splitWithCommand(tc.in)
		if cmd != tc.outCmd {
			t.Fatalf("%d expect %s get %s", i, tc.outCmd, cmd)
		}

		for j, arg := range tc.outArgs {
			if arg != args[j] {
				t.Fatalf("%d expect %s get %s", i, arg, args[j])
			}
		}
	}
}

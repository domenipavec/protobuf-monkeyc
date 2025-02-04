package xtesting

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type MonkeyRunnerCB func(string)

type MonkeyRunnerCase struct {
	InitCode string
	Function string
	Args     []any
	Callback MonkeyRunnerCB
}

type MonkeyRunner struct {
	wg    sync.WaitGroup
	cases []MonkeyRunnerCase
}

func (m *MonkeyRunner) Run(cs MonkeyRunnerCase) {
	m.cases = append(m.cases, cs)
}

func (m *MonkeyRunner) Execute(t *testing.T) {
	casesCode := ""
	caseCodeTmpl := `
	{
		%s
		var got = %s(%s);
		logger.debug(got);
	}
`
	for _, cs := range m.cases {
		t.Log(cs)
		casesCode += fmt.Sprintf(caseCodeTmpl, cs.InitCode, cs.Function, formatArgs(cs.Args))
	}

	code := fmt.Sprintf(`
import Toybox.Lang;
import Toybox.Test;

(:test)
function funcTest(logger as Logger) as Boolean {
%s

	return true;
}
`, casesCode)
	t.Log(code)

	codeFn := filepath.Join(t.TempDir(), "Test.mc")
	err := os.WriteFile(codeFn, []byte(code), 0644)
	require.NoError(t, err)

	pwd, err := os.Getwd()
	require.NoError(t, err)

	dockerArgs := []string{
		"run",
		"--rm",
		"--net", "host",
		"--entrypoint", "/bin/bash",
		"-v", filepath.Dir(pwd) + ":/app",
		"-v", getCertPath(t) + ":/key",
		"-v", codeFn + ":/app/Example/Test.mc",
		"--tmpfs", "/out",
		"-w", "/app/Example",
		"ghcr.io/matco/connectiq-tester:latest",
		"-c", "monkeyc -f 'monkey.jungle;barrels.jungle' -d enduro3 -o /out/test.prg -l 3 -y /key --unit-test && (monkeydo /out/test.prg enduro3 -t || true)",
	}
	t.Log(strings.Join(dockerArgs, " "))
	cmd := exec.Command(
		"docker", dockerArgs...,
	)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	require.NoError(t, err)

	t.Log(string(out))

	i := 0
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "DEBUG") {
			parts := strings.SplitN(line, ": ", 2)
			require.Len(t, parts, 2)

			m.cases[i].Callback(parts[1])
			i++
			if i >= len(m.cases) {
				break
			}
		}
	}
	require.Equal(t, len(m.cases), i, "missing output from some cases")
}

func formatArgs(args []any) string {
	argStr := ""
	for i, a := range args {
		if i != 0 {
			argStr += ", "
		}
		argStr += fmt.Sprint(a)
	}
	return argStr
}

func getCertPath(t *testing.T) string {
	home, err := os.UserHomeDir()
	require.NoError(t, err)

	// TODO: create if doesn't exist
	return filepath.Join(home, "garmin", "developerkey", "developer_key")
}

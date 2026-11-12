package progress

import (
	"io"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func Test_terminalFd_pseudoTerminal(t *testing.T) {
	tty := openPseudoTerminal(t)

	fd, ok := terminalFd(tty)
	assert.True(t, ok)
	assert.Equal(t, int(tty.Fd()), fd)

	_, ok = terminalFd(struct{ io.Writer }{tty})
	assert.False(t, ok)
}

func TestProgress_initForRender_terminal(t *testing.T) {
	tty := openPseudoTerminal(t)
	const cols = 60
	winsize := &unix.Winsize{Row: 24, Col: cols}
	require.NoError(t, unix.IoctlSetWinsize(int(tty.Fd()), unix.TIOCSWINSZ, winsize))

	p := &Progress{}
	p.SetOutputWriter(tty)
	p.initForRender()
	defer p.Stop()

	assert.Eventually(t, func() bool { return p.getTerminalWidth() == cols }, time.Second, 10*time.Millisecond)

	// An explicit width still wins over the width measured from the terminal.
	p.SetTerminalWidth(42)
	assert.Equal(t, 42, p.getTerminalWidth())
}

// openPseudoTerminal returns the terminal side of a new pseudo-terminal. Both
// sides are closed when the test ends.
func openPseudoTerminal(t *testing.T) *os.File {
	t.Helper()

	pty, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		t.Skipf("pseudo-terminal unavailable: %v", err)
	}
	t.Cleanup(func() { _ = pty.Close() })

	// x/sys has no grantpt/unlockpt/ptsname on Linux, and this module does not
	// depend on a pty library, so unlock and name the terminal with ioctls.
	require.NoError(t, unix.IoctlSetPointerInt(int(pty.Fd()), unix.TIOCSPTLCK, 0))
	n, err := unix.IoctlGetInt(int(pty.Fd()), unix.TIOCGPTN)
	require.NoError(t, err)

	tty, err := os.OpenFile("/dev/pts/"+strconv.Itoa(n), os.O_RDWR|unix.O_NOCTTY, 0)
	require.NoError(t, err)
	t.Cleanup(func() { _ = tty.Close() })

	return tty
}

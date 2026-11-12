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

func Test_terminalConn_pseudoTerminal(t *testing.T) {
	tty := openPseudoTerminal(t)

	conn, ok := terminalConn(tty)
	assert.True(t, ok)
	assert.NotNil(t, conn)

	_, ok = terminalConn(struct{ io.Writer }{tty})
	assert.False(t, ok)
}

func TestProgress_initForRender_terminal(t *testing.T) {
	tty := openPseudoTerminal(t)
	const cols = 60
	withFd(t, tty, func(fd int) {
		winsize := &unix.Winsize{Row: 24, Col: cols}
		require.NoError(t, unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ, winsize))
	})
	wasNonBlocking := isNonBlocking(t, tty)

	p := &Progress{}
	p.SetOutputWriter(tty)
	p.initForRender()
	defer p.Stop()

	assert.Eventually(t, func() bool { return p.getTerminalWidth() == cols }, time.Second, 10*time.Millisecond)

	// The file belongs to the caller; measuring it must not change its mode.
	assert.Equal(t, wasNonBlocking, isNonBlocking(t, tty))

	// An explicit width still wins over the width measured from the terminal.
	p.SetTerminalWidth(42)
	assert.Equal(t, 42, p.getTerminalWidth())
}

func TestProgress_updateTerminalSize_closedTerminal(t *testing.T) {
	tty := openPseudoTerminal(t)
	withFd(t, tty, func(fd int) {
		winsize := &unix.Winsize{Row: 24, Col: 60}
		require.NoError(t, unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ, winsize))
	})
	conn, ok := terminalConn(tty)
	require.True(t, ok)

	p := &Progress{}
	p.updateTerminalSize(conn)
	assert.Equal(t, 60, p.getTerminalWidth())

	// the width is unknown once the terminal is gone
	require.NoError(t, tty.Close())
	p.updateTerminalSize(conn)
	assert.Equal(t, 0, p.getTerminalWidth())
}

// isNonBlocking returns true if the file is in non-blocking mode.
func isNonBlocking(t *testing.T, f *os.File) bool {
	t.Helper()

	nonBlocking := false
	withFd(t, f, func(fd int) {
		flags, err := unix.FcntlInt(uintptr(fd), unix.F_GETFL, 0)
		require.NoError(t, err)
		nonBlocking = flags&unix.O_NONBLOCK != 0
	})
	return nonBlocking
}

// withFd calls fn with the file descriptor of f, without the side effect that
// File.Fd has of putting the file into blocking mode.
func withFd(t *testing.T, f *os.File, fn func(fd int)) {
	t.Helper()

	conn, err := f.SyscallConn()
	require.NoError(t, err)
	require.NoError(t, conn.Control(func(fd uintptr) { fn(int(fd)) }))
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

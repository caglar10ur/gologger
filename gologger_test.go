package logger

import "testing"
import "os"

func TestLogger(t *testing.T) {
	generate := func(ll *Logger) {
		ll.Debugln("foo")
		ll.Infoln("foo")
		ll.Warnln("foo")
		ll.Errorln("foo")
		ll.Fatalln("foo")
		ll.Debugf("%.2f %s", 3.14159265359, "bar")
		ll.Infof("%.2f %s", 3.14159265359, "bar")
		ll.Warnf("%.2f %s", 3.14159265359, "bar")
		ll.Errorf("%.2f %s", 3.14159265359, "bar")
		ll.Fatalf("%.2f %s", 3.14159265359, "bar")
	}

	logfile, _ := os.OpenFile("foo", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)

	logger := New(logfile)
	logger.SetLogLevel(Error)
	generate(logger)

	logger.SetFlags(LtraceFlags)
	generate(logger)
	logger.SetFlags(LdebugFlags)
	generate(logger)
	logger.SetFlags(LstdFlags)
	generate(logger)

	logger = New(nil)
	logger.SetLogLevel(Error)
	generate(logger)

	logger.SetFlags(LtraceFlags)
	generate(logger)
	logger.SetFlags(LdebugFlags)
	generate(logger)
	logger.SetFlags(LstdFlags)
	generate(logger)
}

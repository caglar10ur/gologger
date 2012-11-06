package logger

import "testing"

func TestLogger(t *testing.T) {
    logger := New()

    generate := func() {
        logger.Debugln("foo")
        logger.Infoln("foo")
        logger.Warnln("foo")
        logger.Errorln("foo")
        logger.Fatalln("foo")
        logger.Debugf("%.2f %s", 3.14159265359, "bar")
        logger.Infof("%.2f %s", 3.14159265359, "bar")
        logger.Warnf("%.2f %s", 3.14159265359, "bar")
        logger.Errorf("%.2f %s", 3.14159265359, "bar")
        logger.Fatalf("%.2f %s", 3.14159265359, "bar")
    }

    logger.SetLogLevel(Debug)
    generate()
    logger.SetLogLevel(Warn)
    logger.SetFlags()
    generate()
    logger.UnsetFlags()
    generate()

    logger.SetOutput("foo")
    generate()
    logger.SetFlags()
    generate()
}

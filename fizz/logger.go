package fizz

import "github.com/sirupsen/logrus"

var Logger = logrus.WithField("component", "fizz")

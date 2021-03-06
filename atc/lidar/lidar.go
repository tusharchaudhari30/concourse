package lidar

import (
	"os"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/lock"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
)

func NewRunner(
	logger lager.Logger,
	clock clock.Clock,
	scanRunner Runner,
	checkRunner Runner,
	runnerInterval time.Duration,
	notifications Notifications,
	lockFactory lock.LockFactory,
	componentFactory db.ComponentFactory,
) ifrit.Runner {
	return grouper.NewParallel(
		os.Interrupt,
		[]grouper.Member{
			{
				Name:   atc.ComponentLidarScanner,
				Runner: NewIntervalRunner(logger, clock, scanRunner, runnerInterval, notifications, atc.ComponentLidarScanner, lockFactory, componentFactory),
			},
			{
				Name:   atc.ComponentLidarChecker,
				Runner: NewIntervalRunner(logger, clock, checkRunner, runnerInterval, notifications, atc.ComponentLidarChecker, lockFactory, componentFactory),
			},
		},
	)
}

func NewCheckerRunner(
	logger lager.Logger,
	clock clock.Clock,
	checkRunner Runner,
	runnerInterval time.Duration,
	notifications Notifications,
	lockFactory lock.LockFactory,
	componentFactory db.ComponentFactory,
) ifrit.Runner {
	return grouper.NewParallel(
		os.Interrupt,
		[]grouper.Member{
			{
				Name:   atc.ComponentLidarChecker,
				Runner: NewIntervalRunner(logger, clock, checkRunner, runnerInterval, notifications, atc.ComponentLidarChecker, lockFactory, componentFactory),
			},
		},
	)
}

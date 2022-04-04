package gocache

import "time"

type sweeper struct {
	Interval time.Duration
	stop     chan bool
}

func (s *sweeper) run(c *myCache) {
	ticker := time.NewTicker(s.Interval)
	for {
		select {
		case <-ticker.C:
			c.deleteExpired()
		case <-s.stop:
			ticker.Stop()
			return
		}
	}
}

func newSweeper(c *myCache, i time.Duration) *sweeper {
	return &sweeper{
		Interval: i,
		stop:     make(chan bool),
	}
}

func (s *sweeper) stopSweeper() {
	s.stop <- true
}

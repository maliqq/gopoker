package server

import (
	"github.com/golang/glog"
	"github.com/rcrowley/go-metrics"
)

import (
	"time"
)

import (
	"gopoker/play"
)

// StatsWorker - collect node stats
type NodeStatsWorker struct {
	*Node

	total    *metrics.StandardGauge
	waiting  *metrics.StandardGauge
	active   *metrics.StandardGauge
	paused   *metrics.StandardGauge
	inactive *metrics.StandardGauge
}

func (w *NodeStatsWorker) Start() {
	w.total = metrics.NewGauge()
	metrics.Register("rooms.total", w.total)

	w.waiting = metrics.NewGauge()
	metrics.Register("rooms.waiting", w.waiting)

	w.active = metrics.NewGauge()
	metrics.Register("rooms.active", w.active)

	w.paused = metrics.NewGauge()
	metrics.Register("rooms.active", w.active)

	w.inactive = metrics.NewGauge()
	metrics.Register("rooms.inactive", w.inactive)

	for {
		time.After(1 * time.Second)
		w.count()
	}
}

func (w *NodeStatsWorker) count() {
	var total, waiting, active, paused, inactive int64

	for _, room := range w.Node.Rooms {
		total++
		switch room.State {
		case play.Waiting:
			waiting++
		case play.Paused:
			paused++
		case play.Active:
			active++
		case play.Closed:
			inactive++
		}
	}

	w.total.Update(total)
	w.waiting.Update(total)
	w.active.Update(active)
	w.paused.Update(paused)
	w.inactive.Update(inactive)

	glog.Infof("[worker] [stats] %d rooms, %d waiting, %d active, %d paused, %d closed\n",
		total, waiting, active, paused, inactive)
}

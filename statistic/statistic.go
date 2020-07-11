package statistic

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

type GorStat struct {
	statName string
	rateMs   int
	latest   int
	mean     int
	max      int
	count    int
}

func NewStat(statName string, rateMs int) *GorStat {
	s := GorStat{
		statName: statName,
		rateMs:   rateMs,
	}
	log.Printf("%s:%s\n", s.statName, "latest,mean,max,count,count/second,gcount")
	go s.reportStats()

	return &s
}

func (s *GorStat) Write(latest int) {
	if latest > s.max {
		s.max = latest
	}
	if latest != 0 {
		s.mean = ((s.mean * s.count) + latest) / (s.count + 1)
	}
	s.latest = latest
	s.count = s.count + 1
}

func (s *GorStat) Reset() {
	s.latest = 0
	s.max = 0
	s.mean = 0
	s.count = 0
}

func (s *GorStat) String() string {
	return fmt.Sprintf(
		"%s:%s,%s,%s,%s,%s,%s\n",
		s.statName,
		strconv.Itoa(s.latest),
		strconv.Itoa(s.mean),
		strconv.Itoa(s.max),
		strconv.Itoa(s.count),
		strconv.Itoa(s.count/(s.rateMs/1000.0)),
		strconv.Itoa(runtime.NumGoroutine()),
	)
}

func (s *GorStat) reportStats() {
	for {
		log.Println(s)
		s.Reset()
		time.Sleep(time.Duration(s.rateMs) * time.Millisecond)
	}
}

type StatisticCollector struct {
	data  map[string]int
	start time.Time
}

func NewStatisticCollector() *StatisticCollector {
	sc := &StatisticCollector{
		data:  make(map[string]int),
		start: time.Now(),
	}
	go sc.consolePrint() // todo: remove after test
	return sc
}

func (sc *StatisticCollector) Incr(name string) uint {
	sc.data[name]++

	return uint(sc.data[name])
}

func (sc *StatisticCollector) Count(name string, count int) int {
	sc.data[name] = count

	return sc.data[name]
}

func (sc *StatisticCollector) SecondsFromStart() uint {
	return uint(time.Since(sc.start))
}

func (sc *StatisticCollector) consolePrint() {
	for {
		log.Println(sc.data)
		time.Sleep(5 * time.Second)
	}
}

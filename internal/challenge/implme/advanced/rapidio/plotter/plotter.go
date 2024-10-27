package plotter

import (
	"sort"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/romangurevitch/concurrencyworkshop/internal/challenge/implme/advanced/rapidio/simulator"
)

type statistic struct {
	interval int
	avg      time.Duration
	p95      time.Duration
	median   time.Duration
}

func Plot(results []simulator.EventResult, filename string) error {
	p := plot.New()

	p.Title.Text = "Responsiveness Graph"
	p.X.Label.Text = "Latency"
	p.Y.Label.Text = "Interval"

	intervalResultMap := map[int][]time.Duration{}

	for _, result := range results {
		var durations []time.Duration
		var ok bool
		if durations, ok = intervalResultMap[result.Interval]; !ok {
			durations = []time.Duration{result.HandledAt.Sub(result.CreatedAt)}
			intervalResultMap[result.Interval] = durations
		}

		intervalResultMap[result.Interval] = append(durations, result.HandledAt.Sub(result.CreatedAt))
	}
	avg := make(plotter.XYs, 0)
	medians := make(plotter.XYs, 0)
	p95 := make(plotter.XYs, 0)
	for _, statistics := range getStats(intervalResultMap) {
		avg = append(avg, plotter.XY{X: float64(statistics.avg), Y: float64(statistics.interval)})
		medians = append(medians, plotter.XY{X: float64(statistics.median), Y: float64(statistics.interval)})
		p95 = append(p95, plotter.XY{X: float64(statistics.p95), Y: float64(statistics.interval)})
	}

	err := plotutil.AddScatters(p, "avg", avg, "p95", p95, "median", medians)
	if err != nil {
		return err
	}

	// Save the plot to a PNG file.
	if err := p.Save(30*vg.Centimeter, 30*vg.Centimeter, filename); err != nil {
		return err
	}

	return nil
}

func getStats(intervalResultMap map[int][]time.Duration) []statistic {
	var result []statistic

	for interval, latencies := range intervalResultMap {
		result = append(result, calculateStats(interval, latencies))
	}
	return result
}

func calculateStats(interval int, latencies []time.Duration) statistic {
	return statistic{
		interval: interval,
		avg:      average(latencies),
		median:   median(latencies),
		p95:      percentile95(latencies),
	}
}

func average(durations []time.Duration) time.Duration {
	var sum time.Duration
	for _, d := range durations {
		sum += d
	}
	return time.Duration(int64(sum) / int64(len(durations)))
}

func median(durations []time.Duration) time.Duration {
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})
	middle := len(durations) / 2
	if len(durations)%2 == 0 {
		return (durations[middle-1] + durations[middle]) / 2
	}
	return durations[middle]
}

func percentile95(durations []time.Duration) time.Duration {
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})
	p95Index := int(0.95 * float64(len(durations)-1))
	return durations[p95Index]
}

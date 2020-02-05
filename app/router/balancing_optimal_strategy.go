// +build !confonly

package router

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/session"
	"v2ray.com/core/transport"
	"v2ray.com/core/transport/pipe"

	"v2ray.com/core/common/task"
	"v2ray.com/core/features/outbound"
)

// OptimalStrategy pick outbound by net speed
type OptimalStrategy struct {
	timeout       time.Duration
	interval      time.Duration
	url           *url.URL
	count         uint32
	obm           outbound.Manager
	tag           string
	tags          []string
	periodic      *task.Periodic
	periodicMutex sync.Mutex
}

// NewOptimalStrategy create new strategy
func NewOptimalStrategy(config *BalancingOptimalStrategyConfig) *OptimalStrategy {
	s := &OptimalStrategy{}
	if config.Timeout == 0 {
		s.timeout = time.Second * 5
	} else {
		s.timeout = time.Millisecond * time.Duration(config.Timeout)
	}
	if config.Interval == 0 {
		s.interval = time.Second * 60 * 10
	} else {
		s.interval = time.Millisecond * time.Duration(config.Interval)
	}
	if config.Url == "" {
		s.url, _ = url.Parse("https://www.google.com")
	} else {
		var err error
		s.url, err = url.Parse(config.Url)
		if err != nil {
			panic(err)
		}
		if s.url.Scheme != "http" && s.url.Scheme != "https" {
			panic("Only http/https url support")
		}
	}
	if config.Count == 0 {
		s.count = 3
	} else {
		s.count = config.Count
	}

	return s
}

// PickOutbound implement BalancingStrategy interface
func (s *OptimalStrategy) PickOutbound(obm outbound.Manager, tags []string) string {
	if len(tags) == 0 {
		panic("0 tags")
	} else if len(tags) == 1 {
		return s.tag
	}

	s.obm = obm
	s.tags = tags

	if s.periodic == nil {
		s.periodicMutex.Lock()
		if s.periodic == nil {
			s.tag = s.tags[0]
			s.periodic = &task.Periodic{
				Interval: s.interval,
				Execute:  s.run,
			}
			go s.periodic.Start()
		}
		s.periodicMutex.Unlock()
	}

	return s.tag
}

type optimalStrategyTestResult struct {
	tag   string
	score float64
}

// periodic execute function
func (s *OptimalStrategy) run() error {
	tags := s.tags
	count := s.count

	results := make([]optimalStrategyTestResult, len(tags))

	var wg sync.WaitGroup
	wg.Add(len(tags))
	for i, tag := range tags {
		result := &results[i]
		result.tag = tag
		go s.testOutboud(tag, result, count, &wg)
	}
	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		// score scores in desc order
		return results[i].score > results[j].score
	})

	s.tag = results[0].tag
	newError(fmt.Sprintf("Balance OptimalStrategy now pick detour [%s](score: %.2f) from %s", results[0].tag, results[0].score, tags)).AtInfo().WriteToLog()

	return nil
}

// Test outbound's network state with multi-round
func (s *OptimalStrategy) testOutboud(tag string, result *optimalStrategyTestResult, count uint32, wg *sync.WaitGroup) {
	// test outbound by fetch url
	defer wg.Done()

	oh := s.obm.GetHandler(tag)
	if oh == nil {
		newError("Wrong OptimalStrategy tag").AtError().WriteToLog()
		return
	}

	scores := make([]float64, count)
	for i := uint32(0); i < count; i++ {
		client := s.buildClient(oh)
		// send http request though this outbound
		req, _ := http.NewRequest("GET", s.url.String(), nil)
		startAt := time.Now()
		resp, err := client.Do(req)
		// use http response speed or time(no http content) as score
		score := 0.0
		if err != nil {
			newError(err).AtError().WriteToLog()
		} else {
			contentSize := 0
			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				contentSize += len(scanner.Bytes())
			}
			finishAt := time.Now()
			if contentSize != 0 {
				score = float64(contentSize) / (float64(finishAt.UnixNano()-startAt.UnixNano()) / float64(time.Second))
			} else {
				// assert http header's Byte size is 100B
				score = 100 / (float64(finishAt.UnixNano()-startAt.UnixNano()) / float64(time.Second))
			}
		}
		scores[i] = score
		// next test round
		client.CloseIdleConnections()
	}

	// calculate average score and end test round
	var minScore float64 = float64(math.MaxInt64)
	var maxScore float64 = float64(math.MinInt64)
	var sumScore float64
	var score float64

	for _, score := range scores {
		if score < minScore {
			minScore = score
		}
		if score > maxScore {
			maxScore = score
		}
		sumScore += score
	}
	if len(scores) < 3 {
		score = sumScore / float64(len(scores))
	} else {
		score = (sumScore - minScore - maxScore) / float64(s.count-2)
	}
	newError(fmt.Sprintf("Balance OptimalStrategy get %s's score: %.2f", tag, score)).AtDebug().WriteToLog()
	result.score = score
}

func (s *OptimalStrategy) buildClient(oh outbound.Handler) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				netDestination, err := net.ParseDestination(fmt.Sprintf("%s:%s", network, addr))
				if err != nil {
					return nil, err
				}

				uplinkReader, uplinkWriter := pipe.New()
				downlinkReader, downlinkWriter := pipe.New()
				ctx = session.ContextWithOutbound(
					ctx,
					&session.Outbound{
						Target: netDestination,
					})
				go oh.Dispatch(ctx, &transport.Link{Reader: uplinkReader, Writer: downlinkWriter})

				return net.NewConnection(net.ConnectionInputMulti(uplinkWriter), net.ConnectionOutputMulti(downlinkReader)), nil
			},
			MaxConnsPerHost:    1,
			MaxIdleConns:       1,
			DisableCompression: true,
			DisableKeepAlives:  true,
		},
		Timeout: s.timeout,
	}
}

package cmd

import (
	"github.com/nsqio/nsq/nsqctl/pkg/nsqd"
	"github.com/nsqio/nsq/nsqctl/pkg/nsqlookupd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"sync"
)

var topics []string

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a topic or set of topics",
	Run: func(cmd *cobra.Command, args []string) {

		lookupClient := &nsqlookupd.Client{
			Url: lookupdHost,
			Cli: &http.Client{},
		}

		topicMap := make(map[string][]string)
		for _, v := range topics {
			p, err := lookupClient.GetProducersForTopic(v)
			if err != nil {
				log.Fatalf("getting producers %v", err)
			}
			if len(p) == 0 {
				log.Fatalln("no producers returned for topic")
			}
			topicMap[v] = p
		}

		var errs []error
		for k, v := range topicMap {
			for i := 0; i < len(v); i++ {
				doneChan, errChan := pauseTopic(k, v)
			Loop:
				for {
					select {
					case e := <-errChan:
						log.Errorf("%v", e)
						errs = append(errs, e)
					case <-doneChan:
						break Loop
					}
				}
			}
		}

		if len(errs) > 0 {
			log.Errorf("%d or more errors encountered.", len(errs))
		}

		log.Infof("pause complete - %d/%d successful", len(topics)-len(errs), len(topics))
	},
}

func pauseTopic(t string, h []string) (<-chan bool, <-chan error) {
	doneChan := make(chan bool)
	errChan := make(chan error)

	// dispatch calls to pause Topic
	wg := &sync.WaitGroup{}
	for _, v := range h {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nsqdClient := &nsqd.Client{
				Url: v,
				Cli: &http.Client{},
			}
			err := nsqdClient.PauseTopic(t)
			if err != nil {
				errChan <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		doneChan <- true
	}()

	return doneChan, errChan
}

func init() {
	rootCmd.AddCommand(pauseCmd)
	pauseCmd.Flags().StringSliceVar(&topics, "topics", []string{}, "one or more comma separated topics.")
	pauseCmd.MarkFlagRequired("topics")
}

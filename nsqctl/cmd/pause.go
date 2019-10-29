package cmd

import (
	"github.com/nsqio/nsq/nsqctl/pkg/nsdlookupd"
	"github.com/nsqio/nsq/nsqctl/pkg/nsqd"
	"github.com/spf13/cobra"
	"net/http"
)

var topics []string

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a topic or set of topics",
	Run: func(cmd *cobra.Command, args []string) {

		lookupClient := &nsdlookupd.Client{
			Url: lookupdHost,
			Cli: &http.Client{},
		}

		topicMap := make(map[string][]string)
		for _, v := range topics {
			p, _ := lookupClient.GetProducersForTopic(v)
			topicMap[v] = p
		}

		// todo: pause each topic across all producers at topicMap[topic]
		for k, v := range topicMap {
			// k = topic
			// v = list of producers (ip:port)
			for i := 0; i < len(v); i++ {
				nsqdClient := &nsqd.Client{
					Url: v[0],
					Cli: &http.Client{},
				}
				err := nsqdClient.PauseTopic(k)
				if err != nil {
					// track error and report at end.
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
	pauseCmd.Flags().StringSliceVar(&topics, "topics", []string{}, "one or more comma separated topics.")
	pauseCmd.MarkFlagRequired("topics")
}

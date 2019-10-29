package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	//homedir "github.com/mitchellh/go-homedir"
	//"github.com/spf13/viper"
)

var (
	cfgFile     string
	lookupdHost string
)

var rootCmd = &cobra.Command{
	Use:   "nsqctl",
	Short: "A CLI interface to NSQ",
	Long:  `A longer description that spans multiple lines and likely contains`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nsqctl.yaml)")
	rootCmd.Flags().StringVar(&lookupdHost, "nsqlookupd-host", "127.0.0.1:4161", "nsqlookupd host in fqdn/ip:port format")
}

//func initConfig() {
//  if cfgFile != "" {
//    viper.SetConfigFile(cfgFile)
//  } else {
//    // Find home directory.
//    home, err := homedir.Dir()
//    if err != nil {
//      fmt.Println(err)
//      os.Exit(1)
//    }
//
//    viper.AddConfigPath(home)
//    viper.SetConfigName(".nsqctl")
//  }
//
//  viper.AutomaticEnv()
//
//  if err := viper.ReadInConfig(); err == nil {
//    fmt.Println("Using config file:", viper.ConfigFileUsed())
//  }
//}

package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	certFile, keyFile, caFile string
	skipTLS                   bool
	broker                    string
)

var (
	log *logrus.Entry
)

func init() {
	log = logrus.NewEntry(logrus.New())
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafkamonkey",
	Short: "This is a tool for manipulating a kafka service -- mainly for creating stressful situations.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viperConfig()

	// viper.SetDefault("broker", "localhost:1002")
	rootCmd.PersistentFlags().StringVar(&broker, "broker", viper.GetString("broker"), "broker host:port")

	viper.SetDefault("service_key_file", "service.key")
	rootCmd.PersistentFlags().StringVar(&keyFile, "service-key-file", viper.GetString("service_key_file"), "broker key file")

	viper.SetDefault("service_cert_file", "service.cert")
	rootCmd.PersistentFlags().StringVar(&certFile, "service-cert-file", viper.GetString("service_cert_file"), "broker cert file")

	viper.SetDefault("ca_file", "ca.pem")
	rootCmd.PersistentFlags().StringVar(&caFile, "ca-file", viper.GetString("ca_file"), "broker ca file")

	rootCmd.PersistentFlags().BoolVar(&skipTLS, "no-tls", false, "don't make a TLS connection")
}

// viperConfig reads in config file and ENV variables if set.
func viperConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search for config file in current and home directory with name ".kafkamonkey" (without extension).
	viper.AddConfigPath(".")
	viper.AddConfigPath(home)
	viper.SetConfigName(".kafkamonkey")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			fmt.Printf("Problem with config file (%v) :%v\n", viper.ConfigFileUsed(), err)
			os.Exit(1)
		}
	}
}

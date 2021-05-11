package configuration

import (
	"flamingo.me/dingo"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/subchen/go-log"
)

type LocalAwsConfiguration struct {
	Environments map[string]EnvironmentConfiguration
}

type SubscriptionConfiguration struct {
	Protocol     string
	EndPoint     string
	TopicArn     string
	QueueName    string
	Raw          bool
	FilterPolicy string
}

type TopicConfiguration struct {
	Name          string
	Subscriptions []SubscriptionConfiguration
}

type QueueConfiguration struct {
	Name                          string
	ReceiveMessageWaitTimeSeconds int
}

type QueueAttributeConfiguration struct {
	VisibilityTimeout             int
	ReceiveMessageWaitTimeSeconds int
}

type RandomLatencyConfiguration struct {
	Min int
	Max int
}

type SQS struct {
	Queues                 []QueueConfiguration
	QueueAttributeDefaults QueueAttributeConfiguration
	RandomLatency          RandomLatencyConfiguration
}

type EnvironmentConfiguration struct {
	Host      string
	Port      string
	SqsPort   string
	SnsPort   string
	Region    string
	AccountID string
	LogToFile bool
	LogFile   string
	SNS       []TopicConfiguration
	SQS       SQS
}

const DefaultEnvironment = "Local"

func (configuration *LocalAwsConfiguration) Load(fileName string, injector *dingo.Injector) error {
	logger, err := injector.GetInstance(log.Logger{})
	if err != nil {
		return err;
	}

	logger.(*log.Logger).Debugf("Load configuration from %s", fileName)
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	err = config.LoadFiles(fileName)
	if err != nil {
		return err
	}

	err = config.BindStruct("", &configuration.Environments)
	if err != nil {
		return err
	}

	config.LoadOSEnv([]string{"AWS_ENVIRONMENT"}, false)
	env := config.String("AWS_ENVIRONMENT", DefaultEnvironment)
	if len(env) == 0 {
		env = DefaultEnvironment
	}

	logger.(*log.Logger).Debugf("AWS environment is '%s'", env)
	injector.Bind(EnvironmentConfiguration{}).To(configuration.Environments[env])
	return nil
}

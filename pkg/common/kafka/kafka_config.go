package kafka

import (
	"github.com/IBM/sarama"
	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
)

func NewKafkaConfig() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_7_2_0
	if config.Config.Kafka.Username != "" && config.Config.Kafka.Password != "" {
		cfg.Net.SASL.Enable = true
		cfg.Net.SASL.User = config.Config.Kafka.Username
		cfg.Net.SASL.Password = config.Config.Kafka.Password
		cfg.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
		cfg.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	}
	return cfg
}

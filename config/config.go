package config

const (
	Endpoint   = "https://api.enotasgw.com.br/v1"
	EndpointV2 = "https://api.enotasgw.com.br/v2"
)

func Configure(apiKey string) Credentials {
	return Credentials{ApiKey: apiKey}
}

package config

const (
	Endpoint   = "https://api.enotasgw.com.br/v1"
	EndpointV2 = "https://api.enotasgw.com.br/v2"
	EndpointV3 = "https://api2.enotasgw.com.br/v3"
)

func Configure(apiKey string) Credentials {
	return Credentials{ApiKey: apiKey}
}

package config

import (
	"bytes"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	LEDGER_SERVER_HOST                          string = "scalar.dl.client.server.host"
	LEDGER_SERVER_PORT                          string = "scalar.dl.client.server.port"
	LEDGER_SERVER_PRIVILEDGED_PORT              string = "scalar.dl.client.server.privileged_port"
	CERT_PATH                                   string = "scalar.dl.client.cert_path"
	CERT_PEM                                    string = "scalar.dl.client.cert_pem"
	CERT_HOLDER_ID                              string = "scalar.dl.client.cert_holder_id"
	CERT_VERSION                                string = "scalar.dl.client.cert_version"
	PRIVATE_KEY_PATH                            string = "scalar.dl.client.private_key_path"
	PRIVATE_KEY_PEM                             string = "scalar.dl.client.private_key_pem"
	TLS_ENABLED                                 string = "scalar.dl.client.tls.enabled"
	TLS_CA_ROOT_CERT_PATH                       string = "scalar.dl.client.tls.ca_root_cert_path"
	TLS_CA_ROOT_CERT_PEM                        string = "scalar.dl.client.tls.ca_root_cert_pem"
	AUTHORIZATION_CREDENTIAL                    string = "scalar.dl.client.authorization.credential"
	CLIENT_MODE                                 string = "scalar.dl.client.mode"
	PROXY_SERVER                                string = "scalar.dl.client.proxy.server"
	AUDITOR_ENABLED                             string = "scalar.dl.client.auditor.enabled"
	AUDITOR_SERVER_HOST                         string = "scalar.dl.client.auditor.host"
	AUDITOR_SERVER_PORT                         string = "scalar.dl.client.auditor.port"
	AUDITOR_SERVER_PRIVILEDGED_PORT             string = "scalar.dl.client.auditor.privileged_port"
	AUDITOR_TLS_ENABLED                         string = "scalar.dl.client.auditor.tls.enabled"
	AUDITOR_TLS_CA_ROOT_CERT_PATH               string = "scalar.dl.client.auditor.tls.ca_root_cert_path"
	AUDITOR_TLS_CA_ROOT_CERT_PEM                string = "scalar.dl.client.auditor.tls.ca_root_cert_pem"
	AUDITOR_LINEARIZABLE_VALIDATION_ENABLED     string = "scalar.dl.client.auditor.linearizable_validation.enabled"
	AUDITOR_LINEARIZABLE_VALIDATION_CONTRACT_ID string = "scalar.dl.client.auditor.linearizable_validation.contract_id"
)

type ClientConfig struct {
	LedgerHost                              string `validate:"required"`
	LedgerPort                              uint16 `validate:"lt=65536"`
	LedgerPrivilegedPort                    uint16 `validate:"lt=65536"`
	CertHolderId                            string `validate:"required"`
	CertVersion                             int
	Cert                                    string `validate:"required"`
	PrivateKey                              string `validate:"required"`
	IsTlsEnabled                            bool
	TlsCaRootCert                           string `validate:"required_if=IsTlsEnabled true"`
	AuthorizationCredential                 string
	ClientMode                              string `validate:"required,oneof=CLIENT INTERMEDIARY"`
	ProxyServer                             string
	IsAuditorEnabled                        bool
	AuditorHost                             string `validate:"required_if=IsAuditorEnabled true"`
	AuditorPort                             uint16 `validate:"lt=65536"`
	AuditorPrivilegedPort                   uint16 `validate:"lt=65536"`
	IsAuditorTlsEnabled                     bool
	AuditorTlsCaRootCert                    string `validate:"required_if=IsAuditorTlsEnabled true"`
	IsAuditorLinearizableValidationEnabled  bool
	AuditorLinearizableValidationContractId string `validate:"required_if=IsAuditorLinearizableValidationEnabled true"`
}

var validate *validator.Validate = validator.New()

func (c *ClientConfig) Validate() error {
	return validate.Struct(c)
}

func NewClientConfigWithDefaultValues() ClientConfig {
	return ClientConfig{
		LedgerHost:                              "localhost",
		LedgerPort:                              50051,
		LedgerPrivilegedPort:                    50052,
		CertVersion:                             1,
		IsTlsEnabled:                            false,
		IsAuditorEnabled:                        false,
		ClientMode:                              "CLIENT",
		AuditorHost:                             "localhost",
		AuditorPort:                             40051,
		AuditorPrivilegedPort:                   40052,
		IsAuditorTlsEnabled:                     false,
		IsAuditorLinearizableValidationEnabled:  false,
		AuditorLinearizableValidationContractId: "validate-ledger",
	}
}

func NewClientConfigFromJavaProperties(javaProperties string) (ClientConfig, error) {
	var v *viper.Viper = viper.New()
	v.SetConfigType("properties")

	return readConfigByViper(v, javaProperties)
}

func NewClientConfigFromJson(json string) (ClientConfig, error) {
	var v *viper.Viper = viper.New()
	v.SetConfigType("json")

	return readConfigByViper(v, json)
}

func readConfigByViper(v *viper.Viper, configInString string) (clientConfig ClientConfig, err error) {
	if err = v.ReadConfig(bytes.NewBuffer([]byte(configInString))); err != nil {
		return
	}

	clientConfig = NewClientConfigWithDefaultValues()

	if v.GetString(LEDGER_SERVER_HOST) != "" {
		clientConfig.LedgerHost = v.GetString(LEDGER_SERVER_HOST)
	}

	if v.GetUint(LEDGER_SERVER_PORT) != 0 {
		clientConfig.LedgerPort = uint16(v.GetUint(LEDGER_SERVER_PORT))
	}

	if v.GetUint(LEDGER_SERVER_PRIVILEDGED_PORT) != 0 {
		clientConfig.LedgerPrivilegedPort = uint16(v.GetUint(LEDGER_SERVER_PRIVILEDGED_PORT))
	}

	var certPath string = v.GetString(CERT_PATH)
	if certBytes, err := ioutil.ReadFile(certPath); err == nil {
		clientConfig.Cert = string(certBytes)
	}

	var certPem string = v.GetString(CERT_PEM)
	if certPem != "" {
		clientConfig.Cert = certPem
	}

	clientConfig.CertHolderId = v.GetString(CERT_HOLDER_ID)

	if v.GetInt(CERT_VERSION) != 0 {
		clientConfig.CertVersion = v.GetInt(CERT_VERSION)
	}

	var privateKeyPath string = v.GetString(PRIVATE_KEY_PATH)
	if privateKeyBytes, err := ioutil.ReadFile(privateKeyPath); err == nil {
		clientConfig.PrivateKey = string(privateKeyBytes)
	}

	var privateKeyPem string = v.GetString(PRIVATE_KEY_PEM)
	if privateKeyPem != "" {
		clientConfig.PrivateKey = privateKeyPem
	}

	clientConfig.IsTlsEnabled = v.GetBool(TLS_ENABLED)

	var tlsCaRootCertPath string = v.GetString(TLS_CA_ROOT_CERT_PATH)
	if tlsCaRootCertBytes, err := ioutil.ReadFile(tlsCaRootCertPath); err == nil {
		clientConfig.TlsCaRootCert = string(tlsCaRootCertBytes)
	}

	var tlsCaRootCertPem string = v.GetString(TLS_CA_ROOT_CERT_PEM)
	if tlsCaRootCertPem != "" {
		clientConfig.TlsCaRootCert = tlsCaRootCertPem
	}

	if v.GetString(AUTHORIZATION_CREDENTIAL) != "" {
		clientConfig.AuthorizationCredential = v.GetString(AUTHORIZATION_CREDENTIAL)
	}

	if v.GetString(CLIENT_MODE) != "" {
		clientConfig.ClientMode = v.GetString(CLIENT_MODE)
	}

	if v.GetString(PROXY_SERVER) != "" {
		clientConfig.ProxyServer = v.GetString(PROXY_SERVER)
	}

	clientConfig.IsAuditorEnabled = v.GetBool(AUDITOR_ENABLED)

	if v.GetString(AUDITOR_SERVER_HOST) != "" {
		clientConfig.AuditorHost = v.GetString(AUDITOR_SERVER_HOST)
	}

	if v.GetUint(AUDITOR_SERVER_PORT) != 0 {
		clientConfig.AuditorPort = uint16(v.GetUint(AUDITOR_SERVER_PORT))
	}

	if v.GetUint(AUDITOR_SERVER_PRIVILEDGED_PORT) != 0 {
		clientConfig.AuditorPrivilegedPort = uint16(v.GetUint(AUDITOR_SERVER_PRIVILEDGED_PORT))
	}

	clientConfig.IsAuditorTlsEnabled = v.GetBool(AUDITOR_TLS_ENABLED)

	var auditorTlsCaRootCertPath string = v.GetString(AUDITOR_TLS_CA_ROOT_CERT_PATH)
	if auditorTlsCaRootCertBytes, err := ioutil.ReadFile(auditorTlsCaRootCertPath); err == nil {
		clientConfig.AuditorTlsCaRootCert = string(auditorTlsCaRootCertBytes)
	}

	var auditorTlsCaRootCertPem string = v.GetString(AUDITOR_TLS_CA_ROOT_CERT_PEM)
	if auditorTlsCaRootCertPem != "" {
		clientConfig.AuditorTlsCaRootCert = auditorTlsCaRootCertPem
	}

	if clientConfig.IsAuditorEnabled {
		clientConfig.IsAuditorLinearizableValidationEnabled = v.GetBool(AUDITOR_LINEARIZABLE_VALIDATION_ENABLED)
		if v.GetString(AUDITOR_LINEARIZABLE_VALIDATION_CONTRACT_ID) != "" {
			clientConfig.AuditorLinearizableValidationContractId = v.GetString(AUDITOR_LINEARIZABLE_VALIDATION_CONTRACT_ID)
		}
	}

	return
}

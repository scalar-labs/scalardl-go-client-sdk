package config

import (
	"bytes"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	ledgerServerHost                        string = "scalar.dl.client.server.host"
	ledgerServerPort                        string = "scalar.dl.client.server.port"
	ledgerServerPriviledgedPort             string = "scalar.dl.client.server.privileged_port"
	certPath                                string = "scalar.dl.client.cert_path"
	certPem                                 string = "scalar.dl.client.cert_pem"
	certHolderID                            string = "scalar.dl.client.cert_holder_id"
	certVersion                             string = "scalar.dl.client.cert_version"
	privateKeyPath                          string = "scalar.dl.client.private_key_path"
	privateKeyPem                           string = "scalar.dl.client.private_key_pem"
	tlsEnabled                              string = "scalar.dl.client.tls.enabled"
	tlsCaRootCertPath                       string = "scalar.dl.client.tls.ca_root_cert_path"
	tlsCaRootCertPem                        string = "scalar.dl.client.tls.ca_root_cert_pem"
	authorizationCredential                 string = "scalar.dl.client.authorization.credential"
	clientMode                              string = "scalar.dl.client.mode"
	proxyServer                             string = "scalar.dl.client.proxy.server"
	auditorEnabled                          string = "scalar.dl.client.auditor.enabled"
	auditorServerHost                       string = "scalar.dl.client.auditor.host"
	auditorServerPort                       string = "scalar.dl.client.auditor.port"
	auditorServerPriviledgedPort            string = "scalar.dl.client.auditor.privileged_port"
	auditorTLSEnabled                       string = "scalar.dl.client.auditor.tls.enabled"
	auditorTLSCaRootCertPath                string = "scalar.dl.client.auditor.tls.ca_root_cert_path"
	auditorTLSCaRootCertPem                 string = "scalar.dl.client.auditor.tls.ca_root_cert_pem"
	auditorLinearizableValidationEnabled    string = "scalar.dl.client.auditor.linearizable_validation.enabled"
	auditorLinearizableValidationContractID string = "scalar.dl.client.auditor.linearizable_validation.contract_id"
)

// ClientConfig defines the structure of the configurations that is used in ClientService.
// We can use NewClientConfigFromJavaProperties to create it from Java Properties,
// or use NewClientConfigFromJSON to create it from JSON.
type ClientConfig struct {
	LedgerHost                              string `validate:"required"`
	LedgerPort                              uint16 `validate:"lt=65536"`
	LedgerPrivilegedPort                    uint16 `validate:"lt=65536"`
	CertHolderID                            string `validate:"required"`
	CertVersion                             int
	Cert                                    string `validate:"required"`
	PrivateKey                              string `validate:"required"`
	IsTLSEnabled                            bool
	TLSCaRootCert                           string `validate:"required_if=IsTLSEnabled true"`
	AuthorizationCredential                 string
	ClientMode                              string `validate:"required,oneof=CLIENT INTERMEDIARY"`
	ProxyServer                             string
	IsAuditorEnabled                        bool
	AuditorHost                             string `validate:"required_if=IsAuditorEnabled true"`
	AuditorPort                             uint16 `validate:"lt=65536"`
	AuditorPrivilegedPort                   uint16 `validate:"lt=65536"`
	IsAuditorTLSEnabled                     bool
	AuditorTLSCaRootCert                    string `validate:"required_if=IsAuditorTLSEnabled true"`
	IsAuditorLinearizableValidationEnabled  bool
	AuditorLinearizableValidationContractID string `validate:"required_if=IsAuditorLinearizableValidationEnabled true"`
}

var validate *validator.Validate = validator.New()

// Validate checks if mandatory fields are assign and well-formatted.
func (c *ClientConfig) Validate() error {
	return validate.Struct(c)
}

// NewClientConfigWithDefaultValues creates ClientConfig instance with following default values.
// ClientConfig{
//		LedgerHost:                              "localhost",
//		LedgerPort:                              50051,
//		LedgerPrivilegedPort:                    50052,
//		CertVersion:                             1,
//		IsTLSEnabled:                            false,
//		IsAuditorEnabled:                        false,
//		ClientMode:                              "CLIENT",
//		AuditorHost:                             "localhost",
//		AuditorPort:                             40051,
//		AuditorPrivilegedPort:                   40052,
//		IsAuditorTLSEnabled:                     false,
//		IsAuditorLinearizableValidationEnabled:  false,
//		AuditorLinearizableValidationContractID: "validate-ledger",
//	}
func NewClientConfigWithDefaultValues() ClientConfig {
	return ClientConfig{
		LedgerHost:                              "localhost",
		LedgerPort:                              50051,
		LedgerPrivilegedPort:                    50052,
		CertVersion:                             1,
		IsTLSEnabled:                            false,
		IsAuditorEnabled:                        false,
		ClientMode:                              "CLIENT",
		AuditorHost:                             "localhost",
		AuditorPort:                             40051,
		AuditorPrivilegedPort:                   40052,
		IsAuditorTLSEnabled:                     false,
		IsAuditorLinearizableValidationEnabled:  false,
		AuditorLinearizableValidationContractID: "validate-ledger",
	}
}

// NewClientConfigFromJavaProperties parses the given Java Properties string to create ClientConfig according to correspoinding properties.
func NewClientConfigFromJavaProperties(javaProperties string) (ClientConfig, error) {
	var v *viper.Viper = viper.New()
	v.SetConfigType("properties")

	return readConfigByViper(v, javaProperties)
}

// NewClientConfigFromJSON parses the given JSON string to create ClientConfig according to correspoinding properties.
func NewClientConfigFromJSON(json string) (ClientConfig, error) {
	var v *viper.Viper = viper.New()
	v.SetConfigType("json")

	return readConfigByViper(v, json)
}

func readConfigByViper(v *viper.Viper, configInString string) (clientConfig ClientConfig, err error) {
	if err = v.ReadConfig(bytes.NewBuffer([]byte(configInString))); err != nil {
		return
	}

	clientConfig = NewClientConfigWithDefaultValues()

	if v.GetString(ledgerServerHost) != "" {
		clientConfig.LedgerHost = v.GetString(ledgerServerHost)
	}

	if v.GetUint(ledgerServerPort) != 0 {
		clientConfig.LedgerPort = uint16(v.GetUint(ledgerServerPort))
	}

	if v.GetUint(ledgerServerPriviledgedPort) != 0 {
		clientConfig.LedgerPrivilegedPort = uint16(v.GetUint(ledgerServerPriviledgedPort))
	}

	var certPath string = v.GetString(certPath)
	if certBytes, err := ioutil.ReadFile(certPath); err == nil {
		clientConfig.Cert = string(certBytes)
	}

	var certPem string = v.GetString(certPem)
	if certPem != "" {
		clientConfig.Cert = certPem
	}

	clientConfig.CertHolderID = v.GetString(certHolderID)

	if v.GetInt(certVersion) != 0 {
		clientConfig.CertVersion = v.GetInt(certVersion)
	}

	var privateKeyPath string = v.GetString(privateKeyPath)
	if privateKeyBytes, err := ioutil.ReadFile(privateKeyPath); err == nil {
		clientConfig.PrivateKey = string(privateKeyBytes)
	}

	var privateKeyPem string = v.GetString(privateKeyPem)
	if privateKeyPem != "" {
		clientConfig.PrivateKey = privateKeyPem
	}

	clientConfig.IsTLSEnabled = v.GetBool(tlsEnabled)

	var path string = v.GetString(tlsCaRootCertPath)
	if tlsCaRootCertBytes, err := ioutil.ReadFile(path); err == nil {
		clientConfig.TLSCaRootCert = string(tlsCaRootCertBytes)
	}

	var pem string = v.GetString(tlsCaRootCertPem)
	if tlsCaRootCertPem != "" {
		clientConfig.TLSCaRootCert = pem
	}

	if v.GetString(authorizationCredential) != "" {
		clientConfig.AuthorizationCredential = v.GetString(authorizationCredential)
	}

	if v.GetString(clientMode) != "" {
		clientConfig.ClientMode = v.GetString(clientMode)
	}

	if v.GetString(proxyServer) != "" {
		clientConfig.ProxyServer = v.GetString(proxyServer)
	}

	clientConfig.IsAuditorEnabled = v.GetBool(auditorEnabled)

	if v.GetString(auditorServerHost) != "" {
		clientConfig.AuditorHost = v.GetString(auditorServerHost)
	}

	if v.GetUint(auditorServerPort) != 0 {
		clientConfig.AuditorPort = uint16(v.GetUint(auditorServerPort))
	}

	if v.GetUint(auditorServerPriviledgedPort) != 0 {
		clientConfig.AuditorPrivilegedPort = uint16(v.GetUint(auditorServerPriviledgedPort))
	}

	clientConfig.IsAuditorTLSEnabled = v.GetBool(auditorTLSEnabled)

	path = v.GetString(auditorTLSCaRootCertPath)
	if auditorTLSCaRootCertBytes, err := ioutil.ReadFile(path); err == nil {
		clientConfig.AuditorTLSCaRootCert = string(auditorTLSCaRootCertBytes)
	}

	pem = v.GetString(auditorTLSCaRootCertPem)
	if pem != "" {
		clientConfig.AuditorTLSCaRootCert = pem
	}

	if clientConfig.IsAuditorEnabled {
		clientConfig.IsAuditorLinearizableValidationEnabled = v.GetBool(auditorLinearizableValidationEnabled)
		if v.GetString(auditorLinearizableValidationContractID) != "" {
			clientConfig.AuditorLinearizableValidationContractID = v.GetString(auditorLinearizableValidationContractID)
		}
	}

	return
}

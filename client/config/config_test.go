package config

import (
	"testing"
)

func Test_NewClientConfigFromJson_WithEmptyObject_ShouldGetDefaultClientConfig(t *testing.T) {
	var c ClientConfig
	var err error

	if c, err = NewClientConfigFromJson("{}"); err != nil {
		t.Errorf("should be able to create default ClientConfig")
	}

	if c != NewClientConfigWithDefaultValues() {
		t.Errorf("ClientConfig should have default values")
	}
}

func Test_NewClientConfigFromJavaProperties_WithCorrectObject_ShouldGetCorrectClientConfig(t *testing.T) {
	var javaProperties = `
scalar.dl.client.server.host=localhost
scalar.dl.client.server.port=80
scalar.dl.client.server.privileged_port=8080
scalar.dl.client.cert_holder_id=foo
scalar.dl.client.cert_version=100
scalar.dl.client.cert_pem=cert_pem
scalar.dl.client.private_key_pem=private_key_pem
scalar.dl.client.tls.enabled=true
scalar.dl.client.tls.ca_root_cert_pem=ca_root_cert_pem
scalar.dl.client.authorization.credential=credential
scalar.dl.client.mode=INTERMEDIARY
scalar.dl.client.proxy.server=127.0.0.1
scalar.dl.client.auditor.enabled=true
scalar.dl.client.auditor.host=localhost
scalar.dl.client.auditor.port=4040
scalar.dl.client.auditor.privileged_port=40400
scalar.dl.client.auditor.linearizable_validation.enabled=true
scalar.dl.client.auditor.linearizable_validation.contract_id=linearizable
`

	var c ClientConfig
	var err error

	if c, err = NewClientConfigFromJavaProperties(javaProperties); err != nil {
		t.Errorf("can't load Java properties %s", javaProperties)
	}

	if c.LedgerHost != "localhost" {
		t.Errorf("LedgerHost is not match")
	}

	if c.LedgerPort != 80 {
		t.Errorf("LedgerPort is not match")
	}

	if c.LedgerPrivilegedPort != 8080 {
		t.Errorf("LedgerPrivilegedPort is not match")
	}

	if c.CertHolderId != "foo" {
		t.Errorf("CertHolderId is not match")
	}

	if c.CertVersion != 100 {
		t.Errorf("CertVersion is not match")
	}

	if c.Cert != "cert_pem" {
		t.Errorf("Cert is not match")
	}

	if c.PrivateKey != "private_key_pem" {
		t.Errorf("PrivateKey is not match")
	}

	if !c.IsTlsEnabled {
		t.Errorf("IsTlsEnabled is not match")
	}

	if c.TlsCaRootCert != "ca_root_cert_pem" {
		t.Errorf("TlsCaRootCert is not match")
	}

	if c.AuthorizationCredential != "credential" {
		t.Errorf("AuthorizationCredential is not match")
	}

	if c.ClientMode != "INTERMEDIARY" {
		t.Errorf("ClientMode is not match")
	}

	if c.ProxyServer != "127.0.0.1" {
		t.Errorf("ProxyServer is not match")
	}

	if !c.IsAuditorEnabled {
		t.Errorf("IsAuditorEnabled is not match")
	}

	if c.AuditorHost != "localhost" {
		t.Errorf("AuditorHost is not match")
	}

	if c.AuditorPort != 4040 {
		t.Errorf("AuditorPort is not match")
	}

	if c.AuditorPrivilegedPort != 40400 {
		t.Errorf("AuditorPrivilegedPort is not match")
	}

	if !c.IsAuditorLinearizableValidationEnabled {
		t.Errorf("IsAuditorLinearizableValidationEnabled is not match")
	}

	if c.AuditorLinearizableValidationContractId != "linearizable" {
		t.Errorf("AuditorLinearizableValidationContractId is not match")
	}
}

func Test_NewClientConfigFromJson_WithCorrectObject_ShouldGetCorrectClientConfig(t *testing.T) {
	var json = `
{
	"scalar.dl.client.server.host": "localhost",
	"scalar.dl.client.server.port": 80,
	"scalar.dl.client.server.privileged_port": 8080,
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.cert_version": 100,
	"scalar.dl.client.cert_pem": "cert_pem",
	"scalar.dl.client.private_key_pem": "private_key_pem",
	"scalar.dl.client.tls.enabled": true,
	"scalar.dl.client.tls.ca_root_cert_pem": "ca_root_cert_pem",
	"scalar.dl.client.authorization.credential": "credential",
	"scalar.dl.client.mode": "INTERMEDIARY",
	"scalar.dl.client.proxy.server": "127.0.0.1",
	"scalar.dl.client.auditor.enabled": true,
	"scalar.dl.client.auditor.host": "localhost",
	"scalar.dl.client.auditor.port": 4040,
	"scalar.dl.client.auditor.privileged_port": 40400,
	"scalar.dl.client.auditor.linearizable_validation.enabled": true,
	"scalar.dl.client.auditor.linearizable_validation.contract_id": "linearizable"
}
`
	var c ClientConfig
	var err error

	if c, err = NewClientConfigFromJson(json); err != nil {
		t.Errorf("can't load JSON %s", json)
	}

	if c.LedgerHost != "localhost" {
		t.Errorf("LedgerHost is not match")
	}

	if c.LedgerPort != 80 {
		t.Errorf("LedgerPort is not match")
	}

	if c.LedgerPrivilegedPort != 8080 {
		t.Errorf("LedgerPrivilegedPort is not match")
	}

	if c.CertHolderId != "foo" {
		t.Errorf("CertHolderId is not match")
	}

	if c.CertVersion != 100 {
		t.Errorf("CertVersion is not match")
	}

	if c.Cert != "cert_pem" {
		t.Errorf("Cert is not match")
	}

	if c.PrivateKey != "private_key_pem" {
		t.Errorf("PrivateKey is not match")
	}

	if !c.IsTlsEnabled {
		t.Errorf("IsTlsEnabled is not match")
	}

	if c.TlsCaRootCert != "ca_root_cert_pem" {
		t.Errorf("TlsCaRootCert is not match")
	}

	if c.AuthorizationCredential != "credential" {
		t.Errorf("AuthorizationCredential is not match")
	}

	if c.ClientMode != "INTERMEDIARY" {
		t.Errorf("ClientMode is not match")
	}

	if c.ProxyServer != "127.0.0.1" {
		t.Errorf("ProxyServer is not match")
	}

	if !c.IsAuditorEnabled {
		t.Errorf("IsAuditorEnabled is not match")
	}

	if c.AuditorHost != "localhost" {
		t.Errorf("AuditorHost is not match")
	}

	if c.AuditorPort != 4040 {
		t.Errorf("AuditorPort is not match")
	}

	if c.AuditorPrivilegedPort != 40400 {
		t.Errorf("AuditorPrivilegedPort is not match")
	}

	if !c.IsAuditorLinearizableValidationEnabled {
		t.Errorf("IsAuditorLinearizableValidationEnabled is not match")
	}

	if c.AuditorLinearizableValidationContractId != "linearizable" {
		t.Errorf("AuditorLinearizableValidationContractId is not match")
	}
}

func Test_NewClientConfigFromJson_WithInvalidJson_ShouldNotBeValidated(t *testing.T) {
	var c ClientConfig
	var err error

	var withoutCertHolderId = `
{
	"scalar.dl.client.cert_pem": "cert_pem",
	"scalar.dl.client.private_key_pem": "private_key_pem"
}
`

	if c, err = NewClientConfigFromJson(withoutCertHolderId); err != nil {
		t.Errorf("can't load JSON %s", withoutCertHolderId)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated without CertHolderId")
	}

	var withoutCertPem = `
{
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.private_key_pem": "private_key_pem"
}
`

	if c, err = NewClientConfigFromJson(withoutCertPem); err != nil {
		t.Errorf("can't load JSON %s", withoutCertPem)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated without Cert")
	}

	var withoutPrivateKeyPem = `
{
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.cert_pem": "cert_pem"
}
`

	if c, err = NewClientConfigFromJson(withoutPrivateKeyPem); err != nil {
		t.Errorf("can't load JSON %s", withoutPrivateKeyPem)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated without PrivateKey")
	}

	var withoutTlsCaRootCert = `
{
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.cert_pem": "cert_pem",
	"scalar.dl.client.private_key_pem": "private_key_pem",
	"scalar.dl.client.tls.enabled": true
}
`

	if c, err = NewClientConfigFromJson(withoutTlsCaRootCert); err != nil {
		t.Errorf("can't load JSON %s", withoutTlsCaRootCert)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated without TlsCaRootCert")
	}

	var withInvalidClientMode = `
{
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.cert_pem": "cert_pem",
	"scalar.dl.client.private_key_pem": "private_key_pem",
	"scalar.dl.client.mode": "invalidmode"
}
`

	if c, err = NewClientConfigFromJson(withInvalidClientMode); err != nil {
		t.Errorf("can't load JSON %s", withInvalidClientMode)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated with invalid ClientMode")
	}

	var withoutAuditorTlsCaRootCert = `
{
	"scalar.dl.client.cert_holder_id": "foo",
	"scalar.dl.client.cert_pem": "cert_pem",
	"scalar.dl.client.private_key_pem": "private_key_pem",
	"scalar.dl.client.auditor.tls.enabled": true
}
`

	if c, err = NewClientConfigFromJson(withoutAuditorTlsCaRootCert); err != nil {
		t.Errorf("can't load JSON %s", withoutAuditorTlsCaRootCert)
	}

	if err = c.Validate(); err == nil {
		t.Errorf("should not be validated without AuditorTlsCaRootCert")
	}
}
package core

import (
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/prometheus/client_golang/prometheus"
)

// SupportClientImpl ...
type SupportClientImpl struct {
	SupportClient    supportiface.SupportAPI
	checkResultCache map[string]*support.TrustedAdvisorCheckResult
}

// SupportClient ...
type SupportClient interface {
	RequestServiceLimitsRefreshLoop()
	RequestResultCacheRefreshLoop()
	DescribeServiceLimitsCheckResult(checkID string) (*support.TrustedAdvisorCheckResult, error)
}

// SupportExporter ...
type SupportExporter struct {
	supportClient SupportClient
	metricsUsed   map[string]*prometheus.Desc
	metricsLimit  map[string]*prometheus.Desc
}

package phonehome

import (
	"context"
	"runtime"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/highlight-run/highlight/backend/model"
	"github.com/highlight-run/highlight/backend/projectpath"
	"github.com/highlight-run/highlight/backend/util"
	"github.com/highlight/highlight/sdk/highlight-go"
)

type UsageType = string

const AdminUsage UsageType = "highlight-admin-usage"
const WorkspaceUsage UsageType = "highlight-workspace-usage"
const AboutYouSpanName = "highlight-about-you"
const HeartbeatSpanName = "highlight-heartbeat"

const BackendSetup = "highlight-backend-setup"
const SessionCount = "highlight-session-count"
const ErrorCount = "highlight-error-count"
const LogCount = "highlight-log-count"
const SessionViewCount = "highlight-session-view-count"
const ErrorViewCount = "highlight-error-view-count"
const LogViewCount = "highlight-log-view-count"

const AboutYouSpanAdminFirstName = "highlight-about-you-admin-first-name"
const AboutYouSpanAdminLastName = "highlight-about-you-admin-last-name"
const AboutYouSpanAdminEmail = "highlight-about-you-admin-email"
const AboutYouSpanReferral = "highlight-about-you-referral"
const AboutYouSpanRole = "highlight-about-you-role"
const AboutYouSpanTeamSize = "highlight-about-you-team-size"
const AboutYouSpanHeardAbout = "highlight-about-you-heard-about"
const HeartbeatInterval = 5 * time.Second
const HighlightProjectID = "1"
const MetricMemTotal = "highlight-mem-total"
const MetricMemUsedPercent = "highlight-mem-used-percent"
const MetricNumCPU = "highlight-num-cpu"
const SpanDeployment = "highlight-phone-home-deployment-id"
const SpanDopplerConfig = "highlight-doppler-config"
const SpanHighlightVersion = "highlight-version"
const SpanOnPrem = "highlight-is-onprem"

func IsOptedOut(_ context.Context) bool {
	return false
}

func GetDefaultAttributes() ([]attribute.KeyValue, error) {
	cfg, err := projectpath.GetConfig()
	if err != nil {
		return nil, err
	}
	if cfg.PhoneHomeDeploymentID == "" {
		cfg.PhoneHomeDeploymentID = util.GenerateRandomString(32)
	}
	if err = projectpath.SaveConfig(cfg); err != nil {
		return nil, err
	}

	return []attribute.KeyValue{
		attribute.String(highlight.TraceTypeAttribute, string(highlight.TraceTypePhoneHome)),
		attribute.String(highlight.ProjectIDAttribute, HighlightProjectID),
		attribute.String(SpanDeployment, cfg.PhoneHomeDeploymentID),
		attribute.String(SpanDopplerConfig, util.DopplerConfig),
		attribute.String(SpanHighlightVersion, util.Version),
		attribute.String(SpanOnPrem, util.OnPrem),
	}, nil
}

func Start(ctx context.Context) error {
	if IsOptedOut(ctx) {
		return nil
	}

	go func() {
		ctx := context.Background()
		for range time.Tick(HeartbeatInterval) {
			vmStat, _ := mem.VirtualMemory()
			highlight.RecordMetric(ctx, MetricNumCPU, float64(runtime.NumCPU()))
			highlight.RecordMetric(ctx, MetricMemUsedPercent, vmStat.UsedPercent)
			highlight.RecordMetric(ctx, MetricMemTotal, float64(vmStat.Total))
			tags, _ := GetDefaultAttributes()
			tags = append(tags,
				attribute.Int(MetricNumCPU, runtime.NumCPU()),
				attribute.Float64(MetricMemUsedPercent, vmStat.UsedPercent),
				attribute.Int64(MetricMemTotal, int64(vmStat.Total)),
			)

			s, _ := highlight.StartTraceWithTimestamp(ctx, HeartbeatSpanName, time.Now(), []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindServer)}, tags...)
			highlight.EndTrace(s)
		}
	}()
	log.WithContext(ctx).Info("started highlight phone home service")
	return nil
}

func ReportAdminAboutYouDetails(ctx context.Context, admin *model.Admin) {
	if IsOptedOut(ctx) {
		return
	}

	tags, _ := GetDefaultAttributes()
	tags = append(tags, attribute.String(AboutYouSpanRole, ptr.ToString(admin.UserDefinedRole)))
	tags = append(tags, attribute.String(AboutYouSpanTeamSize, ptr.ToString(admin.UserDefinedTeamSize)))
	tags = append(tags, attribute.String(AboutYouSpanHeardAbout, ptr.ToString(admin.HeardAbout)))
	tags = append(tags, attribute.String(AboutYouSpanReferral, ptr.ToString(admin.Referral)))
	if ptr.ToBool(admin.PhoneHomeContactAllowed) {
		tags = append(tags, attribute.String(AboutYouSpanAdminFirstName, ptr.ToString(admin.FirstName)))
		tags = append(tags, attribute.String(AboutYouSpanAdminLastName, ptr.ToString(admin.LastName)))
		tags = append(tags, attribute.String(AboutYouSpanAdminEmail, ptr.ToString(admin.Email)))
	}

	s, _ := highlight.StartTraceWithTimestamp(ctx, AboutYouSpanName, time.Now(), []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindServer)}, tags...)
	highlight.EndTrace(s)
}

func ReportUsageMetrics(ctx context.Context, usageType UsageType, id int, metrics []attribute.KeyValue) {
	if IsOptedOut(ctx) {
		return
	}

	tags, _ := GetDefaultAttributes()
	tags = append(tags, attribute.Int("id", id))
	tags = append(tags, attribute.String("usageType", usageType))
	tags = append(tags, metrics...)
	s, _ := highlight.StartTraceWithTimestamp(ctx, usageType, time.Now(), []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindServer)}, tags...)
	highlight.EndTrace(s)
}

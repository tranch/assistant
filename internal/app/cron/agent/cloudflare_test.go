package agent

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"reflect"
	"testing"

	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
)

func TestDomainAnalyticsReport(t *testing.T) {
	type args struct {
		comp rulebot.IComponent
	}
	tests := []struct {
		name string
		args args
		want []result.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DomainAnalyticsReport(context.Background(), tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DomainAnalyticsReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

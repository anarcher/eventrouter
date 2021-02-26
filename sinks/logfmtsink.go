package sinks

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-logfmt/logfmt"
	"github.com/golang/glog"
	v1 "k8s.io/api/core/v1"
)

type LogfmtSink struct {
}

func NewLogfmtSink() EventSinkInterface {
	return &LogfmtSink{}
}

func (s *LogfmtSink) UpdateEvents(eNew *v1.Event, eOld *v1.Event) {

	var obj string
	obj = fmt.Sprintf("%s/%s", strings.ToLower(eNew.InvolvedObject.Kind), eNew.InvolvedObject.Name)

	rec := []interface{}{
		"namespace", eNew.Namespace,
		"type", eNew.Type,
		"reason", eNew.Reason,
		"object", obj,
		"message", eNew.Message,
		"count", eNew.Count,
	}
	e := logfmt.NewEncoder(os.Stdout)
	s.check(e.EncodeKeyvals(rec...))
	s.check(e.EndRecord())
}

func (s *LogfmtSink) check(err error) {
	if err != nil {
		glog.Errorf("Failed to logfmt error: %v", err)
	}
}

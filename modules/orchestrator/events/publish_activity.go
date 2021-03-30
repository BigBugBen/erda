package events

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
)

type ActivityPublisher struct {
	manager *EventManager
}

func NewActivityPublisher(manager *EventManager) *EventListener {
	var l EventListener = &ActivityPublisher{manager: manager}
	return &l
}

func (p *ActivityPublisher) OnEvent(event *RuntimeEvent) {
	if err := p.manager.publishActivity(event); err != nil {
		logrus.Errorf("[alert] failed to publish activity, event: %v, err is %v", event, err.Error())
	}
}

func (m *EventManager) publishActivity(event *RuntimeEvent) error {
	var action ActionName
	switch event.EventName {
	case RuntimeCreated:
		action = R_ADD
	case RuntimeDeleting:
		action = R_DEL
	case RuntimeDeployStart:
		action = R_DEPLOY_START
	case RuntimeDeployFailed:
		action = R_DEPLOY_FAIL
	case RuntimeDeployCanceling:
		action = R_DEPLOY_CANCEL
	case RuntimeDeployOk:
		action = R_DEPLOY_OK
	default:
		return nil
	}
	content := map[string]interface{}{
		"type":     "R", // Runtime
		"action":   string(action),
		"operator": event.Operator,
		"desc":     "",
		// targets
		"org_id":         strconv.FormatUint(event.Runtime.OrgID, 10),
		"project_id":     strconv.FormatUint(event.Runtime.ProjectID, 10),
		"application_id": strconv.FormatUint(event.Runtime.ApplicationID, 10),
		"runtime_id":     strconv.FormatUint(event.Runtime.ID, 10),
	}
	ctx_, err := json.Marshal(event)
	if err != nil {
		return errors.Wrapf(err, "marshal activity context failed, event: %v, err is %v", event, err.Error())
	}
	ctx := string(ctx_)
	content["context"] = ctx

	message := apistructs.MessageCreateRequest{
		Sender: "orchestrator",
		Labels: map[apistructs.MessageLabel]interface{}{
			apistructs.MySQLLabel: "ps_activities",
		},
		Content: content,
	}
	return m.bdl.CreateMessage(&message)
}

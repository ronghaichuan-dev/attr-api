package attr

import (
	"context"
	"encoding/json"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service/appleapi"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sAttr) SaveNotificationEvent(ctx context.Context, payload string, body *appleapi.ResponseBodyV2) error {
	count, err := dao.AttrAppleNotificationEvent.Ctx(ctx).Where("notification_uuid", body.NotificationUUID).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		marshal, marshalErr := json.Marshal(body)
		if marshalErr != nil {
			return marshalErr
		}
		data := &entity.AttrAppleNotificationEvent{
			Version:          body.Version,
			NotificationUuid: body.NotificationUUID,
			SignedPayload:    payload,
			NotificationType: body.NotificationType,
			Subtype:          body.Subtype,
			ResponseText:     string(marshal),
			ReceivedAt:       time.Now().Unix(),
		}
		if body.Data != nil && body.Data.TransactionInfo != nil {
			if body.Data.Environment != "" {
				data.Envirment = body.Data.Environment
			}
			if body.Data.TransactionInfo.Environment != "" {
				data.Envirment = body.Data.TransactionInfo.Environment
			}
			data.OriginalTransactionId = body.Data.TransactionInfo.OriginalTransactionID
			data.TransactionId = body.Data.TransactionInfo.TransactionID
		}
		_, err = dao.AttrAppleNotificationEvent.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}
		return nil
	}

	return gerror.Newf("[%s]记录已存在", body.NotificationUUID)
}

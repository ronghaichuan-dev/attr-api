package attr

import (
	"context"
	"god-help-service/api/v1/app"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service/appleapi"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sAttr) HandleTrialFreeSubscribe(ctx context.Context, environment, appid, uuid string, payload *app.EventParam, resp *appleapi.ResponseBodyV2) error {
	var count int
	var err error

	if payload != nil {
		count, err = dao.AttrAppSubscriptions.Ctx(ctx).Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, payload.OriginalTransactionId).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			subscription := entity.AttrAppSubscriptions{
				Environment:           environment,
				OrignialTransactionId: payload.OriginalTransactionId,
				Appid:                 appid,
				Uuid:                  uuid,
				IsTrial:               consts.IsTrialFreeYes,
				IsPaid:                consts.IsPaidNo,
				LastEventAt:           time.Now().Unix(),
			}
			_, err = dao.AttrAppSubscriptions.Ctx(ctx).Data(subscription).Insert()
			if err != nil {
				return err
			}
		} else {
			updateFiledData := g.Map{}
			updateFiledData["uuid"] = uuid
			_, err = dao.AttrAppSubscriptions.Ctx(ctx).Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, payload.OriginalTransactionId).Data(updateFiledData).Update()
			if err != nil {
				return err
			}
		}
	}
	if resp != nil {
		count, err = dao.AttrAppSubscriptions.Ctx(ctx).Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, resp.Data.TransactionInfo.OriginalTransactionID).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			subscription := entity.AttrAppSubscriptions{IsTrial: consts.IsTrialFreeYes, IsPaid: consts.IsPaidNo}
			if resp.Data != nil {
				if resp.Data.TransactionInfo != nil {
					subscription.OrignialTransactionId = resp.Data.TransactionInfo.OriginalTransactionID
					subscription.ProductId = resp.Data.TransactionInfo.ProductID
				}
				if resp.Data.AppMetadata != nil {
					subscription.Environment = resp.Data.AppMetadata.Environment
					subscription.Appid = resp.Data.AppMetadata.BundleID
					subscription.Status = resp.Data.AppMetadata.Status
				}
				if resp.Data.RenewalInfo != nil {
					subscription.AutoRenewStatus = resp.Data.RenewalInfo.AutoRenewStatus
				}
				_, err = dao.AttrAppSubscriptions.Ctx(ctx).Data(subscription).Insert()
				if err != nil {
					return err
				}
			}
		} else {
			updateFiledData := g.Map{}
			if resp.Data != nil && resp.Data.TransactionInfo != nil {
				if resp.Data.AppMetadata != nil {
					updateFiledData["environment"] = resp.Data.AppMetadata.Environment
					updateFiledData["status"] = resp.Data.AppMetadata.Status
				}
				updateFiledData["product_id"] = resp.Data.TransactionInfo.ProductID
				if resp.Data.RenewalInfo != nil {
					updateFiledData["auto_renew_status"] = resp.Data.RenewalInfo.AutoRenewStatus
				}
				_, err = dao.AttrAppSubscriptions.Ctx(ctx).Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, resp.Data.TransactionInfo.OriginalTransactionID).Data(updateFiledData).Update()
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

func (s *sAttr) UpdateAppSubscriptionFieldByOriginalTransactionId(ctx context.Context, tx gdb.TX, originalTransactionId string, updateData g.Map) (int64, error) {
	mod := dao.AttrAppSubscriptions.Ctx(ctx)
	if tx != nil {
		mod = mod.TX(tx)
	}
	result, err := mod.Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, originalTransactionId).Data(updateData).Update()
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func (s *sAttr) HandleSubscribe(ctx context.Context, environment, appid, uuid string, payload *app.Payload) error {
	count, err := dao.AttrAppSubscriptions.Ctx(ctx).Where("original_transaction_id", payload.Params.OriginalTransactionId).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		subscription := entity.AttrAppSubscriptions{IsTrial: consts.IsTrialFreeYes, IsPaid: consts.IsPaidNo}
		if payload.Params.OriginalTransactionId != "" {
			subscription = entity.AttrAppSubscriptions{
				Environment:           environment,
				OrignialTransactionId: payload.Params.OriginalTransactionId,
				Appid:                 appid,
				IsTrial:               consts.IsTrialFreeYes,
				IsPaid:                consts.IsPaidNo,
				LastEventAt:           time.Now().Unix(),
			}
		}
		_, err = dao.AttrAppSubscriptions.Ctx(ctx).Data(subscription).Insert()
		if err != nil {
			return err
		}
	} else {
		updateFiledData := g.Map{}
		if payload.Params.OriginalTransactionId != "" {
			updateFiledData["uuid"] = uuid
		}
		_, err = dao.AttrAppSubscriptions.Ctx(ctx).Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, payload.Params.OriginalTransactionId).Data(updateFiledData).Update()
		if err != nil {
			return err
		}
	}

	return nil
}

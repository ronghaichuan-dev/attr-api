package attr

import (
	"context"
	"database/sql"
	"god-help-service/internal/consts"
	"god-help-service/internal/dao"
	"god-help-service/internal/model/entity"
	"god-help-service/internal/service/appleapi"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sAttr) SaveAttrSubscriptionTransaction(ctx context.Context, notificationType string, v2 *appleapi.ResponseBodyV2) error {
	transactionType := ""
	if notificationType == consts.NotificationTypeOneTimeCharge {
		transactionType = "TRIAL"
	} else if notificationType == consts.NotificationTypeRefund {
		transactionType = "REFUND"
	} else if notificationType == consts.NotificationTypeDidRenew {
		transactionType = "RENEW"
	}

	data := entity.AttrSubscriptionTransaction{
		TransactionType: transactionType,
		AppVersion:      v2.Version,
		CreatedAt:       time.Now().Unix(),
	}
	if v2.AppMetadata != nil {
		data.Envirment = v2.AppMetadata.Environment
		data.Appid = v2.AppMetadata.BundleID
	}
	if v2.Data != nil && v2.Data.TransactionInfo != nil {
		var info entity.AttrAppSubscriptions
		err := dao.AttrAppSubscriptions.Ctx(ctx).Fields("uuid").Where(dao.AttrAppSubscriptions.Columns().OrignialTransactionId, v2.Data.TransactionInfo.OriginalTransactionID).Scan(&info)
		if err != nil && !gerror.Equal(err, sql.ErrNoRows) {
			return err
		}
		data.Rsid = info.Rsid

		if v2.Data.TransactionInfo.Environment != "" {
			data.Envirment = v2.Data.TransactionInfo.Environment
		}
		data.OriginalTransactionId = v2.Data.TransactionInfo.OriginalTransactionID
		data.TransactionId = v2.Data.TransactionInfo.TransactionID
		data.InAppOwnership = v2.Data.TransactionInfo.InAppOwnershipType
		data.ProductId = v2.Data.TransactionInfo.ProductID
		data.Price = v2.Data.TransactionInfo.Price
		data.Currency = v2.Data.TransactionInfo.Currency
		data.PurchaseAt = v2.Data.TransactionInfo.PurchaseDate
		if v2.Data.TransactionInfo.BundleID != "" {
			data.Appid = v2.Data.TransactionInfo.BundleID
		}
		_, err = dao.AttrSubscriptionTransaction.Ctx(ctx).Data(data).Insert()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sAttr) UpdateAttrSubscriptionTransaction(ctx context.Context, tx gdb.TX, originalTxId string, updateData gdb.Map) (int64, error) {
	mod := dao.AttrSubscriptionTransaction.Ctx(ctx)
	if tx != nil {
		mod = mod.TX(tx)
	}
	result, err := mod.Where(dao.AttrSubscriptionTransaction.Columns().OriginalTransactionId, originalTxId).Data(updateData).Update()
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

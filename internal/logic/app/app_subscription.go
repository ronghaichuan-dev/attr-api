package app

import (
	"context"
	"god-help-service/internal/dao"
	"god-help-service/internal/service"

	"god-help-service/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppSubscriptions 操作`app_subscriptions`表的DAO结构
type sAppSubscriptions struct {
}

func init() {
	service.RegisterAppSubscriptions(NewAppSubscription())
}

func NewAppSubscription() *sAppSubscriptions {
	return &sAppSubscriptions{}
}

// GetAppSubscriptionByDate 根据AppId和日期获取应用订阅统计信息
func (d *sAppSubscriptions) GetAppSubscriptionByDate(ctx context.Context, appId string, date string) (*entity.AppSubscriptions, error) {
	var appSubscription *entity.AppSubscriptions
	err := dao.AttrAppSubscriptions.Ctx(ctx).Where("app_id", appId).Where("subscription_date", date).Where("deleted_at IS NULL").Scan(&appSubscription)
	return appSubscription, err
}

// UpdateAppSubscription 更新应用订阅统计信息
func (d *sAppSubscriptions) UpdateAppSubscription(ctx context.Context, id int64, count int64, amount float64) error {
	data := g.Map{
		"subscription_count":  count,
		"subscription_amount": amount,
		"updated_at":          gdb.Raw("NOW()"),
	}
	_, err := dao.AttrAppSubscriptions.Ctx(ctx).Data(data).Where("id", id).Where("deleted_at IS NULL").Update()
	return err
}

// CreateTable 创建应用订阅统计表（如果不存在）
func (d *sAppSubscriptions) CreateTable(ctx context.Context) error {
	tableSQL := `
        CREATE TABLE IF NOT EXISTS app_subscriptions (
            id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
            app_id VARCHAR(255) NOT NULL COMMENT '应用ID',
            subscription_count INT DEFAULT 0 COMMENT '订阅总量',
            subscription_amount DECIMAL(10,2) DEFAULT 0.00 COMMENT '订阅总额',
            subscription_date VARCHAR(10) NOT NULL COMMENT '订阅日期（只保存年月日）',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
            deleted_at DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',
            UNIQUE INDEX idx_app_date (app_id, subscription_date)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用订阅统计表';
        `

	_, err := g.DB().Exec(ctx, tableSQL)
	return err
}

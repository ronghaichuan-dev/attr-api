-- ============================================================
-- Migration 002: 收入分析 & 精准投放
-- ============================================================

-- 1. attr_subscription_transaction 新增归因关联字段
ALTER TABLE attr_subscription_transaction
  ADD COLUMN country VARCHAR(32) DEFAULT '' COMMENT '国家',
  ADD COLUMN tracker_network VARCHAR(64) DEFAULT '' COMMENT '归因渠道',
  ADD COLUMN campaign_id VARCHAR(128) DEFAULT '' COMMENT '推广活动ID',
  ADD COLUMN adgroup_id VARCHAR(128) DEFAULT '' COMMENT '广告组ID',
  ADD COLUMN ad_id VARCHAR(128) DEFAULT '' COMMENT '广告ID';

-- 2. 每日聚合统计表
CREATE TABLE IF NOT EXISTS attr_daily_stats (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  stat_date VARCHAR(10) NOT NULL COMMENT '统计日期 YYYY-MM-DD',
  app_id VARCHAR(64) NOT NULL DEFAULT '' COMMENT '应用ID',
  country VARCHAR(32) NOT NULL DEFAULT '' COMMENT '国家',
  tracker_network VARCHAR(64) NOT NULL DEFAULT '' COMMENT '归因渠道',
  campaign_id VARCHAR(128) NOT NULL DEFAULT '' COMMENT '推广活动ID',
  install_count INT DEFAULT 0 COMMENT '安装量',
  trial_count INT DEFAULT 0 COMMENT '试用量',
  subscribe_count INT DEFAULT 0 COMMENT '订阅量（付费）',
  renew_count INT DEFAULT 0 COMMENT '续订量',
  refund_count INT DEFAULT 0 COMMENT '退款量',
  revenue BIGINT DEFAULT 0 COMMENT '收入（分）',
  refund_amount BIGINT DEFAULT 0 COMMENT '退款金额（分）',
  net_revenue BIGINT DEFAULT 0 COMMENT '净收入（分）',
  install_to_trial_rate DECIMAL(5,2) DEFAULT 0 COMMENT '安装转试用率%',
  trial_to_paid_rate DECIMAL(5,2) DEFAULT 0 COMMENT '试用转付费率%',
  created_at BIGINT DEFAULT 0,
  updated_at BIGINT DEFAULT 0,
  UNIQUE KEY uk_date_app_country_network_campaign (stat_date, app_id, country, tracker_network, campaign_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='每日聚合统计';

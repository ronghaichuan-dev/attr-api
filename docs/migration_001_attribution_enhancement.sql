-- ============================================================
-- 归因系统增强 Migration
-- ============================================================

-- 1. 新增 attr_click 点击/展示记录表
CREATE TABLE IF NOT EXISTS `attr_click` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `click_uuid` varchar(64) NOT NULL DEFAULT '' COMMENT '点击唯一ID',
    `app_id` varchar(64) NOT NULL DEFAULT '' COMMENT '应用ID',
    `click_type` varchar(20) NOT NULL DEFAULT 'click' COMMENT '类型: click/impression',
    `idfa` varchar(64) NOT NULL DEFAULT '' COMMENT 'IDFA',
    `idfv` varchar(64) NOT NULL DEFAULT '' COMMENT 'IDFV',
    `gps_adid` varchar(64) NOT NULL DEFAULT '' COMMENT 'GAID',
    `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
    `user_agent` varchar(512) NOT NULL DEFAULT '' COMMENT 'UA',
    `network` varchar(64) NOT NULL DEFAULT '' COMMENT '渠道名称',
    `campaign_id` varchar(128) NOT NULL DEFAULT '' COMMENT '推广活动ID',
    `campaign_name` varchar(256) NOT NULL DEFAULT '' COMMENT '推广活动名称',
    `adgroup_id` varchar(128) NOT NULL DEFAULT '' COMMENT '广告组ID',
    `ad_id` varchar(128) NOT NULL DEFAULT '' COMMENT '广告ID',
    `keyword_id` varchar(128) NOT NULL DEFAULT '' COMMENT '关键词ID',
    `creative` varchar(256) NOT NULL DEFAULT '' COMMENT '素材',
    `click_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '点击链接',
    `redirect_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '跳转地址',
    `click_at` bigint(20) NOT NULL DEFAULT 0 COMMENT '点击时间',
    `created_at` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_click_uuid` (`click_uuid`),
    KEY `idx_app_id` (`app_id`),
    KEY `idx_idfa` (`idfa`),
    KEY `idx_gps_adid` (`gps_adid`),
    KEY `idx_click_at` (`click_at`),
    KEY `idx_network` (`network`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告点击/展示记录表';

-- 2. 新增 attr_postback 回传记录表
CREATE TABLE IF NOT EXISTS `attr_postback` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `app_id` varchar(64) NOT NULL DEFAULT '' COMMENT '应用ID',
    `postback_type` varchar(32) NOT NULL DEFAULT '' COMMENT '回传类型: install/event/reengagement',
    `network` varchar(64) NOT NULL DEFAULT '' COMMENT '渠道',
    `original_transaction_id` varchar(128) NOT NULL DEFAULT '' COMMENT '原始交易ID',
    `event_name` varchar(64) NOT NULL DEFAULT '' COMMENT '事件名',
    `postback_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '回传URL',
    `response_code` int(11) NOT NULL DEFAULT 0 COMMENT '响应码',
    `response_body` text COMMENT '响应内容',
    `status` int(11) NOT NULL DEFAULT 1 COMMENT '状态: 1-成功 2-失败 3-重试中',
    `retry_count` int(11) NOT NULL DEFAULT 0 COMMENT '重试次数',
    `created_at` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_app_id` (`app_id`),
    KEY `idx_network` (`network`),
    KEY `idx_status` (`status`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='归因回传记录表';

-- 3. 完善 attr_device 表 — 新增归因关联字段
ALTER TABLE `attr_device`
    ADD COLUMN `tracker_network` varchar(64) NOT NULL DEFAULT '' COMMENT '归因渠道' AFTER `country`,
    ADD COLUMN `campaign_id` varchar(128) NOT NULL DEFAULT '' COMMENT '推广活动ID' AFTER `tracker_network`,
    ADD COLUMN `adgroup_id` varchar(128) NOT NULL DEFAULT '' COMMENT '广告组ID' AFTER `campaign_id`,
    ADD COLUMN `ad_id` varchar(128) NOT NULL DEFAULT '' COMMENT '广告ID' AFTER `adgroup_id`,
    ADD COLUMN `keyword_id` varchar(128) NOT NULL DEFAULT '' COMMENT '关键词ID' AFTER `ad_id`,
    ADD COLUMN `attr_install_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '关联安装归因记录ID' AFTER `keyword_id`,
    ADD COLUMN `is_first_install` int(11) NOT NULL DEFAULT 0 COMMENT '是否首次安装' AFTER `attr_install_id`,
    ADD COLUMN `channel` varchar(64) NOT NULL DEFAULT '' COMMENT '渠道来源' AFTER `is_first_install`;

-- 4. 完善 attr_install 表 — 新增归因匹配字段
ALTER TABLE `attr_install`
    ADD COLUMN `match_type` varchar(32) NOT NULL DEFAULT '' COMMENT '匹配方式: device_id/referrer/probabilistic/tracker/ad_services' AFTER `token_response_text`,
    ADD COLUMN `match_confidence` varchar(16) NOT NULL DEFAULT '' COMMENT '匹配置信度: high/medium/low' AFTER `match_type`,
    ADD COLUMN `click_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '关联的点击记录ID' AFTER `match_confidence`,
    ADD COLUMN `click_to_install` bigint(20) NOT NULL DEFAULT 0 COMMENT '点击到安装的时间间隔（秒）' AFTER `click_id`,
    ADD COLUMN `ip` varchar(64) NOT NULL DEFAULT '' COMMENT '安装时IP' AFTER `click_to_install`,
    ADD COLUMN `user_agent` varchar(512) NOT NULL DEFAULT '' COMMENT '安装时UA' AFTER `ip`;

-- 5. 完善 attr_app_subscriptions 表 — 新增字段
ALTER TABLE `attr_app_subscriptions`
    ADD COLUMN `offer_type` varchar(32) NOT NULL DEFAULT '' COMMENT '优惠类型' AFTER `expires_at`,
    ADD COLUMN `offer_id` varchar(128) NOT NULL DEFAULT '' COMMENT '优惠ID' AFTER `offer_type`,
    ADD COLUMN `revocation_date` bigint(20) NOT NULL DEFAULT 0 COMMENT '撤销时间' AFTER `offer_id`,
    ADD COLUMN `revocation_reason` int(11) NOT NULL DEFAULT 0 COMMENT '撤销原因' AFTER `revocation_date`;

# TrackFlow v1 — 产品需求文档

## 1. 产品概述

TrackFlow 是一个 Go 微服务，用于 iOS App 用户获客归因和渠道 LTV（Lifetime Value）计算。

**核心能力：**
- 接收客户端设备归因上报，记录用户来源渠道
- 接收客户端设备-订阅关联请求，建立设备与 Apple 订阅的映射
- 接收并处理 Apple App Store Server Notifications V2，跟踪订阅全生命周期
- 按渠道聚合计算 LTV 指标

**目标用户：** iOS App 开发者/增长团队，需要了解各获客渠道的投资回报。

---

## 2. 业务背景

iOS App 的用户可能来自不同渠道（自然流量、Apple Search Ads 等）。为了评估各渠道的获客质量，需要：

1. 在用户首次打开 App 时，记录其来源渠道（归因）
2. 当用户产生订阅行为时，将订阅与设备（渠道）关联
3. 持续跟踪订阅的续订、退款、过期等事件
4. 基于以上数据，按渠道计算 LTV

---

## 3. 功能需求

### 3.1 设备归因上报

**端点：** `POST /v1/devices/attribution`

**触发场景：** App 首次启动或重新安装时，客户端上报归因数据。

**请求参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device_uuid` | UUID string | 是 | 设备唯一标识（IDFV） |
| `channel` | string | 是 | 渠道标识，限定值：`organic`、`apple_search_ads` |
| `campaign_id` | string | 否 | ASA 广告系列 ID |
| `campaign_name` | string | 否 | ASA 广告系列名称 |
| `adgroup_id` | string | 否 | ASA 广告组 ID |
| `adgroup_name` | string | 否 | ASA 广告组名称 |
| `keyword` | string | 否 | ASA 搜索关键词 |
| `region` | string | 否 | 用户地区 |

**业务规则：**

| 规则 | 说明 |
|------|------|
| 首次注册 | 创建设备记录，记录 channel 和 ASA 归因字段 |
| 重复上报（同 device_uuid） | 标记 `is_reinstall=true`，更新 `last_seen_at`，**不覆盖 channel** |
| 输入校验 | `device_uuid` 必须是合法 UUID；`channel` 必须在白名单内 |

**响应：**
- 成功：`201 Created`，返回确认信息（`{"code":201,"message":"ok","data":{"device_uuid":"...","channel":"..."}}`）。注意：此时数据可能尚未入库（异步处理），返回的是请求中的参数回显
- 参数错误：`400 Bad Request`
- 认证失败：`401 Unauthorized`

---

### 3.2 设备-订阅关联

**端点：** `POST /v1/devices/subscription`

**触发场景：** 用户完成订阅购买后，客户端上报设备与订阅的关联关系。

**请求参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `device_uuid` | UUID string | 是 | 设备唯一标识 |
| `original_transaction_id` | string | 是 | Apple 订阅原始交易 ID |

**业务规则：**

| 规则 | 说明 |
|------|------|
| 设备必须已注册 | Handler 同步查询 DB 校验设备是否存在，不存在则返回 404。由于设备归因为异步入库，客户端需确保归因上报已完成后再调用此接口（建议归因上报后延迟 1-2 秒或收到 201 后轮询确认） |
| 新订阅 | 创建 subscription 记录，channel 从关联设备继承 |
| 已有订阅（channel=unknown） | 补偿回填：将 channel 更新为当前设备的 channel（修复 Apple 通知先于客户端上报的时序问题） |
| 已有订阅（channel 已确定） | channel 不可变（ON CONFLICT DO NOTHING），保留首次设备的 channel |
| 换手机场景 | 同一 original_transaction_id 可关联多个设备，但 subscription 的 channel 始终为首次关联设备的 channel |
| 设备-订阅链接 | 多对多关系，重复关联不报错（ON CONFLICT DO NOTHING） |

**响应：**
- 成功：`201 Created`，返回确认信息（`{"code":201,"message":"ok","data":{"device_uuid":"...","original_transaction_id":"..."}}`）。注意：订阅创建/关联为异步处理，返回的是请求参数回显
- 设备不存在：`404 Not Found`（设备校验为同步操作，直接查询 DB）
- 参数错误：`400 Bad Request`
- 认证失败：`401 Unauthorized`

---

### 3.3 Apple 订阅通知处理

**端点：** `POST /v1/apple/notifications`

**触发场景：** Apple App Store Server 推送 V2 通知。

**请求参数：**

| 字段 | 类型 | 说明 |
|------|------|------|
| `signedPayload` | string | Apple JWS 签名载荷 |

**JWS 验证流程：**

1. 解析 JWT header，提取 x5c 证书链
2. 验证证书链至 Apple Root CA G3
3. 使用叶子证书的 ECDSA 公钥验证签名
4. 解码 payload → `NotificationPayload`
5. 解码内部 `signedTransactionInfo` → `TransactionInfo`
6. 解码内部 `signedRenewalInfo` → `RenewalInfo`

**通知类型与订阅状态流转：**

| 通知类型 | 子类型 | 状态变更 | 收入影响 |
|----------|--------|----------|----------|
| `SUBSCRIBED` | `INITIAL_BUY` / `RESUBSCRIBE` | → `active` | `+price` |
| `DID_RENEW` | — | → `active`, `renewal_count+1` | `+price` |
| `EXPIRED` | `VOLUNTARY` / `BILLING_RETRY` / `PRICE_INCREASE` | → `expired` | — |
| `GRACE_PERIOD_EXPIRED` | — | → `expired` | — |
| `REFUND` | — | `total_refund+amount`，若有 revocation → `revoked` | `+refund` |
| `REFUND_REVERSED` | — | `total_refund-amount`（不低于 0） | `-refund` |
| `REVOKE` | — | → `revoked` | — |
| `DID_CHANGE_RENEWAL_STATUS` | `AUTO_RENEW_ENABLED` / `DISABLED` | 更新 `auto_renew_status` | — |
| `DID_FAIL_TO_RENEW` | `GRACE_PERIOD` | → `grace_period` | — |
| `DID_FAIL_TO_RENEW` | `BILLING_RETRY_PERIOD` | → `billing_retry` | — |
| `DID_CHANGE_RENEWAL_PREF` | `DOWNGRADE` / `UPGRADE` | 更新 `auto_renew_product_id` | — |
| `OFFER_REDEEMED` | — | → `active` | `+price` |
| `RENEWAL_EXTENDED` | — | 更新 `expires_date` | — |
| `TEST` | — | 忽略（仅日志） | — |

**关键规则：**

| 规则 | 说明 |
|------|------|
| 幂等性 | 通过 `notification_uuid` UNIQUE 约束保证，重复通知不重复处理 |
| 始终返回 200 | JWS 验证通过后，无论事件落库或入队是否成功，均返回 200。若落库失败则记录错误日志，依赖 Apple 重试机制（Apple 在未收到 200 时会重试）。仅当 JWS 验证失败时也返回 200（避免泄露验证细节） |
| 金额单位 | milliunits（毫单位），如 $9.99 → 9990 |

**Consumer 业务规则（异步处理阶段）：**

| 规则 | 说明 |
|------|------|
| 订阅不存在（SUBSCRIBED/INITIAL_BUY） | 自动创建 subscription 记录，channel 设为 `unknown`，后续客户端调用 POST /v1/devices/subscription 时通过补偿逻辑回填 channel |
| 订阅不存在（其他通知类型） | 跳过状态更新，将事件标记为 `processed`（日志 warning） |
| REFUND_REVERSED 与收入关系 | 仅减少 `total_refund_milliunits`（不低于 0），不调整 `total_revenue_milliunits`。净收入通过 `revenue - refund` 计算，refund 减少即 net_revenue 自然回升 |

**响应：**
- 始终 `200 OK`（Apple 要求）

---

### 3.4 渠道 LTV 查询

**端点：** `GET /v1/ltv/channels`

**触发场景：** 运营/增长团队查询各渠道的 LTV 数据。

**查询参数：**

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| `start_date` | date (YYYY-MM-DD) | 否 | 设备安装起始日期（基于 `devices.first_seen_at`），用于 cohort 筛选 |
| `end_date` | date (YYYY-MM-DD) | 否 | 设备安装截止日期 |

不传日期参数时返回全量聚合数据。传入日期时，仅统计该时间段内安装的设备及其关联订阅的 LTV。

**返回指标：**

| 指标 | 说明 |
|------|------|
| `channel` | 渠道名称 |
| `device_count` | 该渠道的设备安装数 |
| `subscription_count` | 该渠道的订阅数 |
| `conversion_rate` | 转化率（订阅数 / 设备数） |
| `total_revenue_milliunits` | 总收入（毫单位） |
| `total_refund_milliunits` | 总退款（毫单位） |
| `net_revenue_milliunits` | 净收入（收入 - 退款） |
| `avg_ltv_milliunits` | 平均 LTV（净收入 / 订阅数） |
| `avg_renewal_count` | 平均续订次数 |

**数据来源：** 直接基于 `subscriptions` 表聚合查询，`total_revenue_milliunits` 在每次通知事件时原子更新。

---

### 3.5 健康检查

**端点：** `GET /health`

**返回：**
```json
{"status": "ok", "database": "ok"}
```
当数据库不可用时：
```json
{"status": "degraded", "database": "unavailable"}
```
HTTP 状态码：200（正常）/ 503（降级）

---

## 4. 数据模型

### 4.1 devices 表

存储设备归因数据，一个 device_uuid 对应一条记录。

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| `id` | BIGSERIAL | PK | 自增主键 |
| `device_uuid` | UUID | UNIQUE, NOT NULL | 设备标识（IDFV） |
| `channel` | VARCHAR(64) | NOT NULL | 获客渠道 |
| `campaign_id` | VARCHAR(255) | | ASA 广告系列 ID |
| `campaign_name` | VARCHAR(255) | | ASA 广告系列名称 |
| `adgroup_id` | VARCHAR(255) | | ASA 广告组 ID |
| `adgroup_name` | VARCHAR(255) | | ASA 广告组名称 |
| `keyword` | VARCHAR(255) | | ASA 搜索关键词 |
| `region` | VARCHAR(10) | | 用户地区 |
| `is_reinstall` | BOOLEAN | NOT NULL, DEFAULT FALSE | 是否重装 |
| `first_seen_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 首次出现 |
| `last_seen_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 最后出现 |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 创建时间 |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 更新时间 |

### 4.2 subscriptions 表

存储订阅状态，一个 original_transaction_id 对应一条记录。每次 Apple 通知时原子更新。

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| `id` | BIGSERIAL | PK | 自增主键 |
| `original_transaction_id` | VARCHAR(255) | UNIQUE, NOT NULL | Apple 原始交易 ID |
| `channel` | VARCHAR(64) | NOT NULL | 归因渠道（从首次关联设备继承） |
| `product_id` | VARCHAR(255) | | 订阅产品 ID |
| `status` | VARCHAR(32) | NOT NULL, DEFAULT 'unknown' | 订阅状态 |
| `auto_renew_status` | BOOLEAN | NOT NULL, DEFAULT TRUE | 自动续订开关 |
| `auto_renew_product_id` | VARCHAR(255) | | 自动续订目标产品 |
| `renewal_count` | INT | NOT NULL, DEFAULT 0 | 续订次数 |
| `total_revenue_milliunits` | BIGINT | NOT NULL, DEFAULT 0 | 累计收入（毫单位） |
| `total_refund_milliunits` | BIGINT | NOT NULL, DEFAULT 0 | 累计退款（毫单位） |
| `currency` | VARCHAR(3) | NOT NULL, DEFAULT 'USD' | 货币 |
| `environment` | VARCHAR(16) | NOT NULL, DEFAULT 'Production' | 环境 |
| `original_purchase_date` | TIMESTAMPTZ | | 首次购买时间 |
| `expires_date` | TIMESTAMPTZ | | 到期时间 |
| `last_event_at` | TIMESTAMPTZ | | 最后事件时间 |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 创建时间 |
| `updated_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 更新时间 |

### 4.3 subscription_events 表

存储每一条 Apple 通知事件，用于审计和回溯。HTTP Handler 在 JWS 验证后同步写入（`status=pending`），Consumer 处理完业务逻辑后更新为 `processed`。

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| `id` | BIGSERIAL | PK | 自增主键 |
| `notification_uuid` | UUID | UNIQUE, NOT NULL | 幂等键 |
| `original_transaction_id` | VARCHAR(255) | NOT NULL | Apple 原始交易 ID |
| `notification_type` | VARCHAR(64) | NOT NULL | 通知类型 |
| `subtype` | VARCHAR(64) | | 通知子类型 |
| `transaction_id` | VARCHAR(255) | | 交易 ID |
| `product_id` | VARCHAR(255) | | 产品 ID |
| `price_milliunits` | BIGINT | NOT NULL, DEFAULT 0 | 金额（毫单位） |
| `currency` | VARCHAR(3) | NOT NULL, DEFAULT 'USD' | 货币 |
| `environment` | VARCHAR(16) | NOT NULL, DEFAULT 'Production' | 环境 |
| `status` | VARCHAR(16) | NOT NULL, DEFAULT 'pending' | 处理状态：`pending`（已落库待处理）、`processed`（业务逻辑已完成）、`failed`（处理失败） |
| `processed_at` | TIMESTAMPTZ | | 业务逻辑处理完成时间 |
| `error_message` | TEXT | | 处理失败时的错误信息 |
| `signed_date` | TIMESTAMPTZ | | 签名时间 |
| `expires_date` | TIMESTAMPTZ | | 到期时间 |
| `raw_payload` | JSONB | | 原始通知 JSON |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 创建时间 |

**补偿机制：** 定时任务扫描 `status='pending' AND created_at < NOW() - INTERVAL '5 minutes'` 的事件，重新发送至 SQS 队列。确保即使 SQS 消息丢失，事件也不会被遗漏。

### 4.4 device_subscriptions 表

设备与订阅的多对多关联表。

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| `id` | BIGSERIAL | PK | 自增主键 |
| `device_uuid` | UUID | NOT NULL | 设备标识 |
| `original_transaction_id` | VARCHAR(255) | NOT NULL | Apple 原始交易 ID |
| `created_at` | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() | 关联时间 |

**约束：** `UNIQUE(device_uuid, original_transaction_id)`

### 4.5 数据库索引

除主键和 UNIQUE 约束自动创建的索引外，需额外创建以下索引：

| 表 | 索引 | 类型 | 用途 |
|------|------|------|------|
| `devices` | `idx_devices_channel` | B-tree(`channel`) | LTV 聚合查询按渠道分组 |
| `devices` | `idx_devices_first_seen_at` | B-tree(`first_seen_at`) | LTV cohort 按时间筛选 |
| `subscriptions` | `idx_subscriptions_channel` | B-tree(`channel`) | LTV 聚合查询按渠道分组 |
| `subscriptions` | `idx_subscriptions_status` | B-tree(`status`) | 按状态筛选订阅 |
| `subscription_events` | `idx_events_original_txn_id` | B-tree(`original_transaction_id`) | 按订阅查询事件历史 |
| `subscription_events` | `idx_events_pending` | B-tree(`status`, `created_at`) WHERE `status='pending'` | 补偿任务扫描未处理事件（partial index） |
| `subscription_events` | `idx_events_created_at` | B-tree(`created_at`) | 按时间范围查询事件 |
| `device_subscriptions` | `idx_ds_original_txn_id` | B-tree(`original_transaction_id`) | 按订阅查找关联设备 |

### 4.6 并发控制策略

| 操作 | 策略 | SQL 模式 |
|------|------|----------|
| 设备归因首次注册 | `INSERT ... ON CONFLICT (device_uuid) DO UPDATE SET is_reinstall=true, last_seen_at=NOW()`，排除 `channel` 字段 | 无锁冲突 |
| 创建订阅 | `INSERT ... ON CONFLICT (original_transaction_id) DO NOTHING` | 无锁冲突 |
| 订阅 channel 补偿回填 | `UPDATE subscriptions SET channel=$1 WHERE original_transaction_id=$2 AND channel='unknown'` | WHERE 条件保证幂等 |
| 设备-订阅关联 | `INSERT ... ON CONFLICT (device_uuid, original_transaction_id) DO NOTHING` | 无锁冲突 |
| 通知事件插入（HTTP Handler 同步） | `INSERT ... ON CONFLICT (notification_uuid) DO NOTHING`，返回是否实际插入 | 幂等去重，未插入则跳过入队 |
| 通知事件状态更新（Consumer） | `UPDATE subscription_events SET status='processed', processed_at=NOW() WHERE id=$1 AND status='pending'` | WHERE 条件防止重复处理 |
| 订阅状态/收入更新 | `UPDATE subscriptions SET renewal_count = renewal_count + 1, total_revenue_milliunits = total_revenue_milliunits + $1 WHERE original_transaction_id = $2` | 原子递增，无 read-modify-write |
| 退款撤销 | `UPDATE subscriptions SET total_refund_milliunits = GREATEST(total_refund_milliunits - $1, 0) WHERE original_transaction_id = $2` | GREATEST 防止负值 |

---

## 5. 消息队列

使用 AWS SQS FIFO 队列解耦事件处理：

| 队列 | MessageGroupId | DeduplicationId | 用途 |
|------|----------------|-----------------|------|
| `trackflow-device-attribution.fifo` | device_uuid | 自定义 | 设备归因事件 |
| `trackflow-subscription-link.fifo` | original_transaction_id | 自定义 | 设备-订阅关联事件 |
| `trackflow-apple-notification.fifo` | original_transaction_id | notification_uuid | Apple 通知事件 |

每个队列配有对应的死信队列（DLQ）。FIFO 队列保证同一 MessageGroupId 的消息按序处理。

### 5.1 架构模式

**设备归因 / 设备-订阅关联（纯异步）：**

```
  客户端 ──→ HTTP Handler ──→ 参数校验 ──→ SQS FIFO ──→ Consumer ──→ DB
                                              ↑ 入队成功即返回 201
```

**Apple 通知（先落库再入队）：**

```
  Apple ──→ HTTP Handler ──→ JWS 验证
               │
               ├─→ INSERT subscription_events（status=pending） ← 先持久化原始事件
               ├─→ 发送 SQS FIFO 消息
               └─→ 返回 200
                         │
                         ▼
               Consumer ──→ 业务逻辑（状态流转、收入累加）
                         ──→ UPDATE subscription_events status=processed
```

**设计原因：** Apple 通知返回 200 后 Apple 不会重试。若仅依赖 SQS 而消息丢失，该通知将永久丢失。先将原始事件落库可保证数据零丢失，即使 SQS 故障也可通过补偿任务重新处理 pending 状态的事件。

设备归因和订阅关联端点无需先落库，因为客户端具备重试能力。

**同步/异步边界：**

| 端点 | HTTP 响应时机 | 数据持久化时机 |
|------|--------------|---------------|
| `POST /v1/devices/attribution` | 消息入队成功即返回 `201` | Consumer 异步写入 DB |
| `POST /v1/devices/subscription` | 消息入队成功即返回 `201` | Consumer 异步写入 DB |
| `POST /v1/apple/notifications` | 事件落库 + 入队成功即返回 `200` | 原始事件同步写入，业务逻辑 Consumer 异步处理 |
| `GET /v1/ltv/channels` | 直接查询 DB 返回 `200` | — |
| `GET /health` | 直接返回 | — |

**注意：** 设备归因和订阅关联接口返回成功时数据尚未入库，存在短暂的最终一致性窗口。客户端应具备重试容错能力。

### 5.2 Consumer 部署

- HTTP Server 和 SQS Consumer 运行在同一个进程中
- 每个队列启动独立的 goroutine 轮询消费
- 优雅关闭时先停止接收新消息，等待进行中的消息处理完成（10s 超时）

### 5.3 死信队列（DLQ）策略

| 配置项 | 值 | 说明 |
|--------|------|------|
| `maxReceiveCount` | 3 | 消息处理失败 3 次后转入 DLQ |
| DLQ 消息保留期 | 14 天 | 提供足够的排查时间 |
| 告警 | DLQ 消息数 > 0 时触发告警 | 通过 CloudWatch Alarm → SNS 通知 |
| 重放 | 提供 CLI 命令将 DLQ 消息重新发送至源队列 | `trackflow dlq redrive --queue <name>` |

---

## 6. 非功能需求

### 6.1 安全

| 项目 | 要求 |
|------|------|
| API 认证 | 客户端接口（3.1、3.2）使用 API Key 认证，通过 `X-API-Key` 请求头传递 |
| Apple 通知端点 | 不做 API Key 认证（Apple 无法携带自定义 header），安全性通过 JWS 签名验证保证 |
| LTV 查询端点 | 使用 API Key 认证，建议分配独立的只读 API Key |
| Apple JWS 验证 | 完整 x5c 证书链验证至 Apple Root CA G3 |
| 无硬编码密钥 | 所有敏感配置通过环境变量注入 |
| SQL 注入防护 | 全部使用参数化查询（pgx） |
| Rate Limiting | 基于 IP 的令牌桶限流 |

**API Key 管理：**
- API Key 存储在环境变量 `TRACKFLOW_API_KEYS` 中，支持多个 Key（逗号分隔）
- 中间件校验：无效或缺失 Key 返回 `401 Unauthorized`
- v1 阶段使用静态 API Key，后续可升级为 JWT 或 OAuth2

### 6.2 可靠性

| 项目 | 要求 |
|------|------|
| 幂等性 | notification_uuid UNIQUE 约束防重 |
| 通知保序 | SQS FIFO MessageGroupId=original_transaction_id |
| 优雅关闭 | SIGINT/SIGTERM → 10s 超时等待请求完成 |
| Apple 通知响应 | JWS 验证通过后始终返回 200（无论后续落库/入队是否成功），避免 Apple 重复推送。落库失败时记录错误日志，依赖 Apple 自身重试 |
| Apple 通知零丢失 | 先落库 subscription_events 再入队 SQS，补偿任务兜底扫描 pending 事件 |

### 6.3 可观测性

| 项目 | 实现 |
|------|------|
| 结构化日志 | slog JSON handler |
| 请求追踪 | X-Request-ID 中间件 |
| 健康检查 | GET /health（检查数据库连接） |
| 业务指标 | Prometheus metrics，通过 `GET /metrics` 端点暴露 |

**Prometheus 指标清单：**

| 指标名 | 类型 | 标签 | 说明 |
|--------|------|------|------|
| `trackflow_http_requests_total` | Counter | `method`, `path`, `status` | HTTP 请求总数 |
| `trackflow_http_request_duration_seconds` | Histogram | `method`, `path` | HTTP 请求延迟 |
| `trackflow_sqs_messages_received_total` | Counter | `queue` | SQS 接收消息数 |
| `trackflow_sqs_messages_processed_total` | Counter | `queue`, `status` | SQS 处理消息数（success/error） |
| `trackflow_apple_notifications_total` | Counter | `type`, `subtype` | Apple 通知按类型计数 |
| `trackflow_subscriptions_active` | Gauge | `channel` | 当前活跃订阅数 |
| `trackflow_dlq_messages` | Gauge | `queue` | DLQ 中的消息数 |

### 6.4 部署

| 项目 | 方案 |
|------|------|
| 容器化 | 多阶段 Docker 构建（Go 1.23 builder + distroless） |
| 本地开发 | docker-compose（PostgreSQL 16 + LocalStack SQS） |
| 数据库迁移 | golang-migrate，支持 up/down |

---

## 7. API 响应格式

### 成功响应

```json
{
  "code": 201,
  "message": "ok",
  "data": { ... }
}
```

### 错误响应

```json
{
  "code": 400,
  "message": "device_uuid is required"
}
```

### 状态码约定

| 状态码 | 场景 |
|--------|------|
| 200 | 查询成功、Apple 通知接收 |
| 201 | 创建/关联请求已受理（异步处理） |
| 400 | 请求参数校验失败 |
| 401 | API Key 缺失或无效 |
| 404 | 资源不存在（如设备未注册） |
| 429 | 请求频率超限 |
| 500 | 服务器内部错误 |
| 503 | 服务降级（数据库不可用） |

---

## 8. 关键业务流程

### 8.1 完整用户生命周期

```
1. 用户安装 App
   → 客户端调用 POST /v1/devices/attribution → 返回 201（已受理）
   → [异步] Consumer 消费 SQS 消息 → 写入 devices 表

2. 用户购买订阅
   → 客户端调用 POST /v1/devices/subscription → 返回 201（已受理）
   → [异步] Consumer 创建 subscription（继承 channel），建立 device-subscription 链接

3. Apple 推送 SUBSCRIBED 通知
   → POST /v1/apple/notifications
   → Handler: JWS 验证 → 写入 subscription_events（pending） → 入队 SQS → 返回 200
   → [异步] Consumer: subscription status=active, revenue+price → 事件标记 processed

4. 每月续订
   → Apple 推送 DID_RENEW → 同上流程
   → [异步] Consumer: renewal_count+1, revenue+price

5. 用户退款
   → Apple 推送 REFUND → 同上流程
   → [异步] Consumer: total_refund+amount

6. 运营查询 LTV
   → GET /v1/ltv/channels（同步查询 DB）
   → 返回各渠道的安装数、订阅数、转化率、净收入、平均 LTV
```

### 8.2 换手机场景

```
1. 用户在设备 A（来自 ASA）购买订阅 → subscription.channel = apple_search_ads
2. 用户换到设备 B（自然流量）→ 客户端再次调用 POST /v1/devices/subscription
3. subscription.channel 不变（仍为 apple_search_ads）
4. device_subscriptions 表新增 设备B-订阅 的链接记录
```

---

## 9. 渠道定义

| 渠道标识 | 说明 |
|----------|------|
| `organic` | 自然流量（App Store 搜索、推荐等） |
| `apple_search_ads` | Apple Search Ads 广告获客 |

后续可扩展其他渠道（需修改 `device.validChannels` 白名单）。

---

## 10. 金额单位说明

所有金额字段使用 **milliunits（毫单位）** 存储，避免浮点精度问题。

| 实际金额 | 存储值 | 字段后缀 |
|----------|--------|----------|
| $9.99 | 9990 | `_milliunits` |
| $0.00 | 0 | `_milliunits` |
| ¥6.00 | 6000 | `_milliunits` |

Apple 通知中的 `price` 字段即为 milliunits 单位。

---

## 11. 环境变量配置

| 变量名 | 必填 | 默认值 | 说明 |
|--------|------|--------|------|
| `PORT` | 否 | `8080` | HTTP 服务端口 |
| `DATABASE_DSN` | 是 | — | PostgreSQL 连接串，如 `postgres://user:pass@host:5432/trackflow?sslmode=disable` |
| `DB_MAX_OPEN_CONNS` | 否 | `25` | 数据库最大连接数 |
| `DB_MAX_IDLE_CONNS` | 否 | `5` | 数据库最大空闲连接数 |
| `DB_CONN_MAX_LIFETIME` | 否 | `5m` | 连接最大存活时间 |
| `SQS_ENDPOINT` | 否 | — | SQS 端点（本地开发填 LocalStack 地址，生产留空使用 AWS 默认） |
| `SQS_QUEUE_PREFIX` | 否 | `trackflow` | SQS 队列名前缀 |
| `AWS_REGION` | 是 | — | AWS 区域 |
| `TRACKFLOW_API_KEYS` | 是 | — | API Key 列表，逗号分隔 |
| `APPLE_ROOT_CA_PATH` | 否 | 内嵌 | Apple Root CA G3 证书路径，不填则使用内嵌证书 |
| `RATE_LIMIT_RPS` | 否 | `100` | 每 IP 每秒最大请求数 |
| `RATE_LIMIT_BURST` | 否 | `200` | 令牌桶突发容量 |
| `LOG_LEVEL` | 否 | `info` | 日志级别：`debug`、`info`、`warn`、`error` |
| `ENVIRONMENT` | 否 | `production` | 运行环境：`development`、`staging`、`production` |

---

## 12. 时序竞态场景与补偿

### 12.1 Apple 通知先于客户端上报

```
时间线：
  T1: 用户购买订阅
  T2: Apple 推送 SUBSCRIBED/INITIAL_BUY 通知 → 服务端收到
  T3: 客户端调用 POST /v1/devices/subscription → 服务端收到

问题：T2 时 subscription 尚不存在（客户端还没上报关联）

处理策略：
  T2: SUBSCRIBED 通知到达 → 自动创建 subscription（channel='unknown'）
      → 记录 subscription_event → 累加 revenue
  T3: 客户端上报关联 → 发现 subscription 已存在但 channel='unknown'
      → 补偿回填 channel 为当前设备的渠道
```

### 12.2 设备归因晚于订阅关联

```
时间线：
  T1: 客户端调用 POST /v1/devices/subscription → 设备不存在 → 返回 404

处理策略：
  客户端需保证调用顺序：先归因上报，后订阅关联。
  客户端收到 404 时应重试（指数退避，最多 3 次）。
```

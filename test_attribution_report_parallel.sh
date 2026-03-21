#!/bin/bash

# 归因上报测试参数
appids=("b4cdb7e4-6b40-4b34-84f3-e1d884fd53ca" "4caf15d5-6883-42c2-8533-5bbbc8eb695d" "08bba532-f9a6-4f1b-9f6c-6c3b65d06dc7")

# 并发数
CONCURRENCY=3

# 生成随机字符串函数
generate_random_string() {
    local length=$1
    openssl rand -base64 32 | tr -dc 'a-zA-Z0-9' | head -c $length
}

# 生成归因测试文件
generate_attribution_test_file() {
    local index=$1
    local timestamp=$(date +%s)
    local random_appid=${appids[$RANDOM % ${#appids[@]}]}
    local random_uuid="user_$(generate_random_string 8)"
    local random_idfa="IDFA_$(generate_random_string 32)"
    local random_idfv="IDFV_$(generate_random_string 32)"
    local random_gps_adid="GPS_$(generate_random_string 16)"
    local random_android_id="ANDROID_$(generate_random_string 16)"
    local countries=("CN" "US" "JP" "GB" "DE")
    local random_country=${countries[$RANDOM % ${#countries[@]}]}
    local os_names=("iOS" "Android")
    local random_os=${os_names[$RANDOM % ${#os_names[@]}]}
    local networks=("apple_search_ads" "google_ads" "facebook" "tiktok")
    local random_network=${networks[$RANDOM % ${#networks[@]}]}
    local channels=("organic" "paid" "referral")
    local random_channel=${channels[$RANDOM % ${#channels[@]}]}

    cat > test_attribution_${index}.json << EOF
{
  "environment": "Sandbox",
  "app_token": "test_token",
  "app_version": "2.0",
  "appid": "$random_appid",
  "uuid": "$random_uuid",
  "idfa": "$random_idfa",
  "idfv": "$random_idfv",
  "gps_adid": "$random_gps_adid",
  "android_id": "$random_android_id",
  "os_name": "$random_os",
  "os_version": "${RANDOM:0:1}.${RANDOM:0:1}.${RANDOM:0:1}",
  "language": "zh-CN",
  "country": "$random_country",
  "tracker": "test_tracker",
  "tracker_token": "tracker_token_$(generate_random_string 8)",
  "tracker_uid": "tracker_uid_$(generate_random_string 8)",
  "tracker_version": "1.0",
  "network": "$random_network",
  "channel": "$random_channel",
  "campaign_id": "campaign_$(generate_random_string 8)",
  "adgroup_id": "adgroup_$(generate_random_string 8)",
  "ad_id": "ad_$(generate_random_string 8)",
  "keyword_id": "keyword_$(generate_random_string 8)",
  "agency": "test_agency",
  "install_at": $timestamp,
  "sent_at": $timestamp,
  "is_handle_token": 0,
  "attr_uuid": $timestamp
}
EOF

    echo "生成归因测试文件 ${index}:
  appid: $random_appid
  uuid: $random_uuid
  country: $random_country
  network: $random_network
  os: $random_os"
}

# 运行单个归因测试实例
run_single_attribution_test() {
    local index=$1
    local file="test_attribution_${index}.json"

    echo "运行归因测试实例 ${index}..."
    /Users/ronghaichuan/go/bin/go-wrk -c 3 -d 3 -M POST \
        -H "Content-Type: application/json" \
        -H "Authorization: Basic cGFzc3dvcmQ6Q3NSc1pKayNma2QyMDQ4Zmo=" \
        -body @$file \
        http://localhost/app/v1/attribution/report > test_attribution_result_${index}.txt 2>&1

    echo "归因测试实例 ${index} 完成"
}

# 运行归因测试
run_attribution_test() {
    echo "生成归因测试文件..."

    # 生成多个测试文件
    for ((i=1; i<=CONCURRENCY; i++)); do
        generate_attribution_test_file $i
    done

    echo "运行并行归因测试..."

    # 并行运行测试
    for ((i=1; i<=CONCURRENCY; i++)); do
        run_single_attribution_test $i &
    done

    # 等待所有测试完成
    wait

    echo "所有归因测试完成"

    # 显示测试结果
    echo "\n归因测试结果汇总:"
    for ((i=1; i<=CONCURRENCY; i++)); do
        echo "\n--- 归因测试实例 ${i} 结果 ---"
        cat test_attribution_result_${i}.txt
    done

    # 清理临时文件
    echo "\n清理临时文件..."
    rm -f test_attribution_*.json test_attribution_result_*.txt
}

# 执行归因测试
run_attribution_test

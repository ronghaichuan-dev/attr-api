#!/bin/bash

# 生成随机测试参数
event_codes=("Install" "Subscribe" "TrialFree" "StartUp" "ScreenUnlock" "SubscribeFix")
appids=("b4cdb7e4-6b40-4b34-84f3-e1d884fd53ca" "4caf15d5-6883-42c2-8533-5bbbc8eb695d" "08bba532-f9a6-4f1b-9f6c-6c3b65d06dc7")

# 并发数
CONCURRENCY=3

# 生成随机字符串函数（修复版本）
generate_random_string() {
    local length=$1
    openssl rand -base64 32 | tr -dc 'a-zA-Z0-9' | head -c $length
}

# 生成动态测试文件
generate_test_file() {
    local index=$1
    local random_code=${event_codes[$RANDOM % ${#event_codes[@]}]}
    local timestamp=$(date +%s)
    local random_appid=${appids[$RANDOM % ${#appids[@]}]}
    local random_uuid="user_$(generate_random_string 8)"
    local random_event_uuid="event_$(generate_random_string 12)_$timestamp"
    local random_transaction_id="tx_$(generate_random_string 10)"

    cat > test_request_${index}.json << EOF
{
  "appid": "$random_appid",
  "event_code": "$random_code",
  "uuid": "$random_uuid",
  "event_uuid": "",
  "environment": "Sandbox",
  "app_token": "test_token",
  "app_version": "2.0",
  "country": "CN",
  "created_at": $timestamp,
  "send_at": $timestamp,
  "params": {
    "originalTransactionId": "$random_transaction_id",
    "transaction_id": "$random_transaction_id"
  }
}
EOF

    echo "生成测试文件 ${index}:"
    echo "  appid: $random_appid"
    echo "  event_code: $random_code"
    echo "  uuid: $random_uuid"
    echo "  event_uuid: $random_event_uuid"
}

# 运行单个测试实例
run_single_test() {
    local index=$1
    local file="test_request_${index}.json"

    echo "运行测试实例 ${index}..."
    /Users/ronghaichuan/go/bin/go-wrk -c 3 -d 3 -M POST \
        -H "Content-Type: application/json" \
        -H "Authorization: Basic cGFzc3dvcmQ6Q3NSc1pKayNma2QyMDQ4Zmo=" \
        -body @$file \
        http://localhost/app/v1/event/report > test_result_${index}.txt 2>&1

    echo "测试实例 ${index} 完成"
}

# 运行测试
run_test() {
    echo "生成测试文件..."

    # 生成多个测试文件
    for ((i=1; i<=CONCURRENCY; i++)); do
        generate_test_file $i
    done

    echo "运行并行测试..."

    # 并行运行测试
    for ((i=1; i<=CONCURRENCY; i++)); do
        run_single_test $i &
    done

    # 等待所有测试完成
    wait

    echo "所有测试完成"

    # 显示测试结果
    echo "\n测试结果汇总:"
    for ((i=1; i<=CONCURRENCY; i++)); do
        echo "\n--- 测试实例 ${i} 结果 ---"
        cat test_result_${i}.txt
    done

    # 清理临时文件
    echo "\n清理临时文件..."
    rm -f test_request_*.json test_result_*.txt
}

# 执行测试
run_test

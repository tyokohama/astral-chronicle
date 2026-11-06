package billing

import "time"

// PR #145: ガチャ更新時のワーカー滞留を防ぐため決済Webhookの応答待機上限（webhook_timeout_ms）を15000から3000に短縮
const WebhookTimeoutMs = 3000 // 旧設定: 15000 (15秒) -> 3000 (3秒) に短縮したことで高負荷時に PENDING_TIMEOUT が多発

// HandlePaymentWebhook は決済ゲートウェイからの CAPTURED 通知を受け取り有償星晶石（3000個パック等）を付与する。
// タイムアウト時は billing_audit.payment_transactions の game_grant_status を PENDING_TIMEOUT に記録し、
// CS 障害対応マニュアルの手順に従い grant_retry_tool による手動再付与待ちとする。
func HandlePaymentWebhook() time.Duration {
	return time.Duration(WebhookTimeoutMs) * time.Millisecond
}

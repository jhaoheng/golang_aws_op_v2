## 文件
- https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/connect

## workshop
- https://repost.aws/knowledge-center/connect-outbound-calls-api
- https://aws.amazon.com/tw/blogs/contact-center/automating-outbound-calling-to-customers-using-amazon-connect/


## 購買電話號碼 (在 Amazon Connect 中購買免費電話號碼（TFN）)
1.	登入 Amazon Connect 管理主控台。
2.	選擇你的 Instance，進入「電話號碼（Phone Numbers）」頁面。
3.	點擊「CLAIM A PHONE NUMBER」（申請電話號碼）。
4.	選擇以下參數：
    •	國家/地區：選擇你想要申請電話號碼的國家或地區。
    •	號碼類型（Type）：選擇 Toll-Free。
5.	系統會列出可用的 Toll-Free Number，選擇一個並點擊「CLAIM」。


# 費用

## aws connect 固定費用 (分鐘)
- 0.018

## 電話號碼
- 直接內線撥號(DID)
    - 分配給特定用戶或部門的電話號碼，用於直接撥通目標分機或部門，而不需要經過總機接線員或自動話務系統。
    - 主要用於 公司內部的通訊需求，幫助用戶直接聯繫到特定的員工或團隊。
    - ex: 每個員工或部門可以有自己的獨立號碼。
- 免付費電話號碼(TFN)
    - 通話費用由企業負擔
    - 只能在本地 (本地通話費)
- 環球國際免費電話號碼(UIFN)
    - 通話費用由企業支付
    - 可跨國 (會有跨國際的通話費)

## 傳入呼叫 (每分鐘)
- 有分成 DID, TFN, UIFN 等費率

## 傳出呼叫 (每分鐘)

## 範例情境: 
1. 客戶打過來給客服
2. 客服接起電話，花了三分鐘對話 => 3 * 0.018 (aws connect 服務用量費用)
3. 客服轉接給指定部門，花了四分鐘 (轉接一開始就計費，包含對話時間) => 4 * 0.018 (aws connect 服務用量費用)
4. 使用 電話號碼 => 1天 * N
5. 傳入呼叫費用 => 7(分鐘) * N
6. 傳出呼叫費用 => 4(分鐘) * N
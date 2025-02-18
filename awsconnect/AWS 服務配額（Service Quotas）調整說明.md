# AWS 服務配額（Service Quotas）調整說明
> https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html

1. 服務配額可以調整
	•	除非特別註明，否則 AWS 服務的 所有配額（Service Quotas） 都可以申請調整。
	•	以前，這些配額被稱為 限制（Limits），但現在 AWS 統一改稱為 配額（Quotas）。

2. 每個 AWS 帳戶都有預設配額
	•	當你使用 AWS 服務時，每個帳戶都會有 預設的服務配額。
	•	這些配額 限制了資源的使用量，例如：
	•	Amazon Connect 的外撥電話數量限制
	•	EC2 執行個體（Instances）數量限制
	•	API 呼叫頻率限制

## 目錄
•	重要須知
•	Amazon Connect 配額
•	AppIntegrations 配額
•	Amazon Q in Connect 配額
•	案件（Cases）配額
•	Contact Lens 配額
•	Customer Profiles 配額
•	外撥行銷活動（Outbound Campaigns）配額
•	Voice ID 配額
•	如何計算聯絡數量
•	功能規格
•	Amazon Connect 呼叫中心預設可撥打的國家
•	API 限流（Throttling）配額

## Important things to know

在請求增加服務配額之前，您必須先建立 Amazon Connect 實例（Instance）。

我們會審核每一個 配額增加請求：
	•	小幅度的增加請求：通常可以在 幾小時內 獲得批准。
	•	大幅度的增加請求：需要時間進行 審核、處理、批准和部署，可能 最長需 3 週。
	•	全球範圍的大型增加請求：可能需要 數個月。

📌 如果您的配額增加是較大專案的一部分，請提前規劃時間，確保項目順利進行。

配額的兩種類型
	1.	帳戶層級配額（Account Level Quotas）
	•	影響該 AWS 帳戶及區域（Region）內的所有 Amazon Connect 實例。
	•	例如：某個 API 每秒最大交易數（TPS，Transactions Per Second） 限制。
	2.	資源層級配額（Resource Level Quotas）
	•	僅影響特定 Amazon Connect 實例內的資源。
	•	例如：每個實例的最大使用者數量。
	•	無法在帳戶層級調整，只能針對單個實例調整。

配額適用範圍
	•	配額適用於每個 AWS 區域（Region），您可以在 同一區域內建立多個 Amazon Connect 實例。
	•	可以申請提高該區域內所有實例的配額。
	•	文件中的預設配額值適用於新帳戶，但隨著時間變化，您帳戶的實際配額值可能低於文件中描述的預設值。

📌 並非所有配額都可以調整，部分配額是固定的。


如何查看與管理資源層級配額
	•	需要 AWS CLI 版本 2.13.20 或以上 才能檢視和管理某些資源層級的配額，例如 每個實例的電話號碼數量。


美國電話號碼攜入（Porting）
	•	若您想將現有的美國電話號碼從目前的電信業者轉移至 Amazon Connect，可以使用 相同的申請表單 提出請求。
	•	了解更多電話號碼攜入資訊，請參閱 將現有電話號碼攜入 Amazon Connect。https://docs.aws.amazon.com/connect/latest/adminguide/port-phone-number.html
# Amazon Connect 預設可撥打的國家

> https://docs.aws.amazon.com/connect/latest/adminguide/country-code-allow-list.html

Amazon Connect 能夠撥打哪些國家，取決於你的 AWS 實例（Instance）所在的區域。

若要查看所有可撥打國家的完整清單，可以參考 Amazon Connect 定價頁面（可能需要登入 AWS 帳戶）。

此外，由於 Amazon 可能會調整服務配額（Service Quotas），你目前允許撥打的國家，可能與官方列出的清單不同。


## 特定國家與地區的撥打限制

某些國家的手機號碼 可能不在預設允許清單內，如果你無法撥打這些號碼，需要提交「服務配額增加請求（Service Quota Increase Request）」：

英國（UK）
	•	以下前綴的英國手機號碼 可能無法直接撥打：
+447
	•	若無法撥打這些號碼，必須提交 服務配額增加請求。

日本（Japan）
	•	以下前綴的日本手機號碼 可能無法直接撥打：
+8170、+8180、+8190
	•	若無法撥打這些號碼，也需要提交 服務配額增加請求。


## 如何申請開通其他國家的通話權限

如果需要允許撥打額外的國家，或限制某些國家的通話，請依照以下步驟操作：
1.	登入 AWS 帳戶
	•	進入 AWS Support Console（支援主控台）。
2.	填寫申請表單
	•	選擇服務（Service）：選擇 Amazon Connect（Number Management）。
	•	選擇類別（Category）：選擇 Country Allowlisting for Outbound Calls（國際撥打國家允許清單）。
	•	選擇嚴重性（Severity）：根據需求選擇優先等級。
3.	提供詳細資訊
	•	主題（Subject）：填入簡單標題，例如「Request to enable outbound calling to [Country]」。
	•	描述（Description）：列出你希望開通（或限制）的國家清單。
4.	提交請求
	•	在「Solve now or contact us」頁面：
	•	選擇 Contact us（聯絡我們）。
	•	選擇偏好的聯絡語言 和 聯絡方式（電子郵件或電話）。
	•	提交請求。
5.	等待 Amazon Connect 團隊回覆
	•	Amazon 會審查你的申請，並回覆是否批准你的請求。


## 總結
•	你預設可撥打的國家，取決於 AWS 實例所在的區域。
•	某些國家的手機號碼（如英國 +447、日本 +8170、+8180、+8190）需要額外授權才能撥打。
•	若要開通其他國家的撥打權限，需要提交 AWS 支援請求，並提供詳細的說明與國家清單。
•	申請提交後，Amazon Connect 團隊會審核並回覆結果。
# CalendEye
### 画像を送るとその内容をGoogleカレンダーに登録してくれるLINEBotです。
![image](https://github.com/claustra01/CalendEye/assets/108509532/a90cba43-c66f-4eff-ad68-108a97591fc5)

## 技術的なお話
- Botの機能については、認証や画像の操作、LineBotなど全ての機能をGoの標準パッケージだけで実装しています。(今はgodotenvとlib/pqのみ使用しています)
- GitHubのRepositoryを整えたり、CI(lint/test/build)を組んだりしました。issueやPRのフォーマットも整備しています。
- interfaceを用いてDIがしっかりした構成を目指しています。(現在作業中)

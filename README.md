# Construction sync

受朋友所託，寫一個能將工程照片分類並複製的程式。

## 開始使用

請按照以下說明安裝並設定 Construction sync。

### 下載

請點擊此 [連結](https://github.com/ZhaoTzuHsien/construction-sync/releases/latest/download/sync.exe) 或前往 [Releases](releases) 頁面下載最新版本的程式。

Windows Defender 有可能會把這個檔案視為病毒，並阻止你下載。
這是由於這個程式沒有經過簽署，且 Windows Defender 沒見過這個檔案才覺得有問題。
請放心，所有的程式碼都是安全的，如果你依舊不放心，Construction sync 是一個開源軟體，你可以自行審閱其中的程式碼。

### 設定 config.yaml

在使用程式之前，你需要設定 config.yaml Construction sync 能順利運作。

你可以在以下任一位置新增 config.yaml：
```
├── 內有 sync.exe 的資料夾
|   ├── sync.exe
│   ├── config.yaml
│   └── configs
│       └── config.yaml
└── %AppData%
    └── construction-sync
        └── config.yaml     <- 推薦位置
```

如果你還是不知道要在哪裡新增 config.yaml，你可以點擊兩下 `sync.exe` 執行程式，接著就會看到錯誤訊息告訴你應該要將 config.yaml 放在哪裡。

config.yaml 的內容應該包含：
```yaml
source:
  path: "要複製的資料夾路徑 Ex: /test/data/B073 施工照片(每日、工地祭祀)"
  glob: "非必須，在 source.path 下檔案資料夾的 glob syntax，預設為 ???年??月/*/*"
destination:
  path: "複製目的地的資料夾路徑 Ex: /test/data/B079 個單元施工紀錄"
```
完整設定請參考 [configs/config.yaml](configs/config.yaml)

## 使用

點擊兩下 `sync.exe` 就可以使用啦，就是這麼簡單！

## 開放原始碼授權

本專案採用 MIT 授權，如需更進一步的資訊，請詳閱 [LICENSE.md](LICENSE.md)。


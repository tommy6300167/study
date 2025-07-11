project structure

main.go → 註冊 router
↳ router.go → 分組路由
↳ handler/ → 每個功能一組檔案
↳ todo.go → 呼叫 service、解析參數
↳ service/（可選）→ 執行邏輯、處理流程
↳ store/ → 實際資料操作（CRUD）
↳ model/ → 結構定義

language: golang
frameword: gin
purpose: create todolist api

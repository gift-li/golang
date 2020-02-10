# GoLang -- Coding space
    ### practice and learn about Go Language
* Project ongoing
    * ATM2.0(Latest version)
[toc]
------
### 更新日誌
* 2020/2/8 ATM1.0 上線
    * 上傳內容
        * 單一帳號、名稱和金錢
        * 使用數字清單選擇指令(離開系統方式不明:待調整)
        * struct定義帳號格式
    * 程式格式
        * functions
            * 各功能funcs
        * main
        * system
            * 清單系統func
    * 待修正
        * 離開系統方式不明:待調整
        * 帳號複數，可由使用者自行輸入、存取
* 2020/2/10 ATM2.0 上線
    * 上傳內容
        * ~~導入interface(提供自定義功能函式，提高維護方便性)~~
            * scan出現bug，後續版本處理該問題
        * 誤觸預防
            * 更改密碼新增是否確定更改
            * 提款金額不得大於帳戶金額
        * 離開系統成功(破壞for迴圈條件而非goto)
    * 程式格式
        * func
            * 初次創帳func
            * 清單func
            * 各功能funcs
        * main
            * 定義acc變數
        * settings
            * interface和struct格式
    * 待修正
        * 帳號複數
            * 想法:
                1. 可用2維矩陣存取，掃描資料時針對特定資料使用
                2. 針對特定類別指定特定參數，再用for一一檢查各層資料的該類別
                * 問題: 無法指定存取特定帳戶的資料(都要重頭找)
        * interface功能存疑
            * 學習不夠
                * 再去網路上找資源學習
                * 找書
            1. scanln和err的用法
                * 使用原因
                    * 避免輸入錯誤或失敗造成panic
                * 問題
                    * 一直panic
                    * ~~雙程式並行~~
                        * 關係不明，可能非interface問題
                    * 使用make(map[int]interface{})製作帳號矩陣
                        * assertion需做好
                        * 再去找資料
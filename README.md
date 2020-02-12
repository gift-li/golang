# GoLang -- Coding space
    practice and learn about Go Language
* Project ongoing
    * ATM
        * 2.0(Latest version)
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
* 2020/2/12 ATM3.0構思發想
    * 資源
        * 1.0 + 2.0的程式庫
    * 功能
        * 原本
            * 改帳戶名稱
            * 改密碼
            * 改金額
            * 離開
        * 添加
            * 可登入登出的系統
                * 啟動程式便進入Log in / 初始啟動設定頁面
            * 帳戶基本資料管理
                * 帳戶擁有者可維護個人資料
            * 進入系統需輸入帳號密碼
                * 可辨識該組資料
    * 規劃
        * main.go
            * 主要使用package其他檔案的func
            * ~~input變數儲存空間~~
                * 不確定，待議
        * func.go
            * 功能func
                * 原先+添加
        * type.go
            * 設定處理account資料的interface 和定義帳戶資料的struct
            * ~~interface的func~~
                * 待議，看要獨立或是放到func.go裡
    * 待解決問題
        * 多個帳號密碼
            * 使用map[string/int]interface{}
            * 使用二維矩陣
                * eg.[x][0]放帳號名稱, [x][0]放密碼
        * 登入登出系統
            * 不確定是否屬於golang的範疇
            * 感覺整體複雜成前端與後端(還是說本來就該如此設計)
            * 若參雜到MySQL等後端，此計畫擱置 / 放棄此發想
        * interface不熟悉
            * 找參考書與網路資源
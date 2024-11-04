資料夾結構
================

# cmd目錄
    main啟動指令，每個不同的啟動指令放在不同資料夾中

# pkg 目錄
    pkg 主要的程式碼目錄

## pkg/blockchain 目錄
    區塊鏈相關程式碼，用interface 來定義各種區塊鏈的接口，區塊鏈跟系統整合起來的地方
### pkg/blockchain/third 目錄
    實作呼叫鏈上api

## pkg/const 目錄
    系統相關的常量定義

## pkg/lib 目錄
    將常用的功能封裝起來，會是最底層，不相依專案中的其他package (包含其他lib)

## pkg/config 目錄
    定義config 設定的內容，只能相依lib

## pkg/helper 目錄
    綜合的使用lib config const目錄，封裝為共用func

## pkg/service 目錄
    由ctl呼叫

## pkg/repo 目錄
    使用model操作資料

## pkg/model 目錄
    定義model結構


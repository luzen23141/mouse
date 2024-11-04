Mouse
================

# 指令

## 啟動api伺服器

```bash
./mouse api 
```

## 產地址

```bash
# ./mouse gen {?鏈}
./mouse gen eth
```

#### 產生指定後綴的地址

```bash
# ./mouse gen {鏈} {?後綴}
./mouse gen sol abcd 
```

#### 指定多個執行緒

```bash
# ./mouse gen {鏈} {?後綴} -c {執行緒數}
./mouse gen tron abcd -c 10
```

#### 返回地址的私鑰(私鑰 非hd wallet) (預設為hd wallet的助記詞)

```bash
# ./mouse gen {鏈} {?後綴} --priv
./mouse gen sol --priv
```

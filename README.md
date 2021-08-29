# Fast LSC Deployer

### ⚠ HIGHLY EXPERIMENTAL

`LINE SMART CITY GovTechプログラム` で公式に開発されているものではありません。  
このアプリケーションを利用して、正しくデプロイできる保証はありません。

## デプロイ

`lsc.sh` があるディレクトリ配下で、

### 並列で全てデプロイ

```bash
fld deploy --all --useContainer
```

### 並列でWebUIと汎用APIをデプロイ

```bash
fld deploy --admin-web --platform --useContainer
```

### 逐次でWebUIと汎用APIをデプロイ

```bash
fld deploy --admin-web --platform --sequential --useContainer
```

### dockerを使わずに逐次でWebUIと汎用APIをデプロイ

```bash
fld deploy --admin-web --platform --sequential
```

## ベンチマーク

数回ベンチマークを取ったわけではなく、1回ずつ実行した結果です。

- OS: Ubuntu 20.04.2 LTS on WSL2
- CPU: AMD Ryzen 7 3700X (8 core, 16 thread)
- RAM: WSL2 16GB, Host 64GB
- SSD: Silicon Power M.2 PCIe P34A80

### lsc.sh

```bash
 ❯ ./lsc.sh deploy --useContainer
[INFO]
AWS Profile Name: lsc-dev

...

[INFO] 権限グループを付与しました
[INFO] LSC環境のデプロイが完了しました。
>>> elapsed time 14m47s
```

### Fast LSC Deployer

```bash
 ❯ fld deploy --all --useContainer
 distribution | [INFO]
 distribution | AWS Profile Name: lsc-dev
 
 ...
 
 distribution | [INFO] Finished deploying SAM.
 distribution | Stack Name: lsc-dev-mohe-n-distribution-dynamic
 INFO   |
 INFO   | ビルド結果
 INFO   | bosai : OK
 INFO   | bi : OK
 INFO   | scenario : OK
 INFO   | survey : OK
 INFO   | platform : OK
 INFO   | liff : OK
 INFO   | distribution : OK
 INFO   | admin-web : OK
 INFO   |
 INFO   | デプロイ結果
 INFO   | bosai : OK
 INFO   | bi : OK
 INFO   | survey : OK
 INFO   | scenario : OK
 INFO   | platform : OK
 INFO   | liff : OK
 INFO   | admin-web : OK
 INFO   | distribution : OK
>>> elapsed time 3m29s
```

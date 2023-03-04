## Wine Cellar

### 初期セットアップ

```bash

wine-cellar init
wine-cellar ver add wine #wineコマンドをwine-cellarに追加します
wine-cellar env add main win64 ~/.wine64 wine #wine64をmainという名前で~/.wine64に登録します
wine-cellar set main #mainを現在のバージョンに設定します
```

### 設定を開く

```bash
wine-cellar # 現在の設定されているEnvirionmentを確認
wine-cellar run winecfg # 現在のバージョンのwinecfgを実行
```

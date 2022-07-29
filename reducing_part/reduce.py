import pandas as pd

dfA = pd.read_csv('sorted.csv',names=('time','lati','long','url','tag'),low_memory=False).astype(str)
n = len(dfA.index)
newdf = pd.DataFrame(columns=['time','lati','long','url', 'tag'])
tmp = dfA['tag'].iloc[0]
# dfAは二つのcsvを結合してtagでソートされている状態
# やるべきこと=> 各tagで100個ずつのデータだけを抜き出す
# 実装方法　cntが100個以下なら追加　それ以外ならスルー(新しいものが出てきた場合はそっちに置き換える)
cnt = 0

for i in range(n):
    if(i % 100000 == 0):
        print(i)    
    now = dfA['tag'].iloc[i] #tmp = 連続するもの, now = 現在のもの
    if(tmp != now):
        tmp = now
        cnt = 0
        newdf = newdf.append(dfA.iloc[i], ignore_index=True)
    else:
        if(cnt < 100):
            newdf = newdf.append(dfA.iloc[i], ignore_index=True)
            cnt += 1
    
newdf.to_csv('tags.csv',index=False)
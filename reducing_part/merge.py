import pandas as pd

dfA = pd.read_csv('geotag.csv',names=('id','time','lati','long','url'),low_memory=False).astype(str)
dfB = pd.read_csv('tag.csv',names=('id','tag'),low_memory=False).astype(str)

df = pd.merge(dfA,dfB,on='id',how='right')
df.to_csv('connect.csv',index=False)
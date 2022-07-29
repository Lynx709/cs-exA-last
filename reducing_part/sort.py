import pandas as pd
df = pd.read_csv('connect.csv', names=('id','time','lati','long','url','tag'), low_memory=False).astype(str)
df = df.sort_values(by=["tag","time"])

df.to_csv('true.csv', index=False)
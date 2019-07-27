import os
import json
import sys
import redis
import pandas as pd
import redis_repo

# Replaces with your configuration information
redis_host = "redis"
redis_port = 6379
redis_password = ""

try:
    r = redis.StrictRedis(host=redis_host,
    port=redis_port,
    password=redis_password,
    decode_responses=True)
except Exception as e:
    print(e)

def add_user_table():
    df = pd.read_csv('./data_user.csv', sep='|', names=['Name', 'Firstname', 'Email', 'Phone'], engine="c")
    for index, row in df.iterrows():
        if index != 0 and r:
            redis_repo.add_hash_set_item(r, 'email:' + row['Email'], 'Name', row['Name'])
            redis_repo.add_hash_set_item(r, 'email:' + row['Email'], 'Firstname', row['Firstname'])
            redis_repo.add_hash_set_item(r, 'email:' + row['Email'], 'Phone', row['Phone'])
            redis_repo.add_hash_set_item(r, 'email:' + row['Email'], 'Email', row['Email'])
            print('Name: ' + row['Name'], 'Firstname: ' + row['Firstname'], 'Email: ' + row['Email'])
            print("Add to database")
    print("End")

add_user_table()

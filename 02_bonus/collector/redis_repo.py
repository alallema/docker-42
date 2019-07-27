# All Functions to put data on redis

import redis

# basics repo functions

# Vector    string
# Value     dict
def add_sorted_set_item(r, key, values):
    # print(type(values))
    r.zadd(key, values)

# Field     string
# Value     dict
def add_hash_set_item(r, key, field, value):
    # print(type(value))
    r.hset(key, field, value)

def add_set_item(r, key, values):
    # print(type(values))
    r.sadd(key, values)

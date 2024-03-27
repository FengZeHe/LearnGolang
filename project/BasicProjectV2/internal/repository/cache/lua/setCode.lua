local key = KEYS[1]
local setKey = "CACHE/users/sms/" .. key
local value = ARGV[1]
redis.call("set", setKey, value)
redis.call("expire", setKey, 600)


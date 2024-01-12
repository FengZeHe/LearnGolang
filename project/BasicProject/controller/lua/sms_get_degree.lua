local key     = KEYS[1]
local countkey = "CACHE/users/smscount/" .. key
local existed = redis.call("get", countkey)
local ttl     = tonumber(redis.call("ttl", countkey))
if existed ~= false then
    local count = redis.call('decr', countkey)
    if count - 1 > 0 then
        return 1
    else
        redis.call("set", countkey, -1)
        redis.call("expire", countkey, ttl)
        return -1
    end
else
    redis.call("set", countkey, 5)
    redis.call("expire", countkey, 600)
    return 1
end

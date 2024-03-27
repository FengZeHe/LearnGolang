local key = KEYS[1]
local countKey = "CACHE/users/smscount/" .. key
local existed = redis.call("get", countKey)
local ttl = tonumber(redis.call("ttl", countKey))

if existed ~= false then
    -- 有发送剩余次数 -1 可以发送
    local count = redis.call('decr', countKey)
    if count - 1 > 0 then
        return 1
    else
        -- 没有剩余次数 设置countKey剩余次数为-1,设置过期时间
        redis.call("set", countKey, -1)
        redis.call("expire", countKey, ttl)
        return -1
    end
else
    -- 设置定期发送的剩余次数
    redis.call("set", countKey, 5)
    redis.call("expire", countKey, 600)
    return 1
end
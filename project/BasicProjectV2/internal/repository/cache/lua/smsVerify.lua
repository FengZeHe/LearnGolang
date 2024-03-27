local key = KEYS[1]
local setKey = "CACHE/users/sms/" .. key
local expectCode = ARGV[1]
local existed = redis.call("get", setKey)
if existed ~= false then
    local verify = redis.call("get", setKey)
    if verify == expectCode then
        redis.call("del", setKey)
        return 1
    else
        return -1
    end
else
    return -1
end
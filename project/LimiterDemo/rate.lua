-- 控制2秒内只发20次请求（平均10次/秒）
local count = 0
local max_requests = 20  -- 10次/秒 × 2秒 = 20次

function request()
    if count < max_requests then
        count = count + 1
        return wrk.format(nil, "/user/12345")
    else
        -- 达到最大请求数后，不再发送新请求
        return nil
    end
end
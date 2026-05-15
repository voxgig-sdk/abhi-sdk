-- Abhi SDK error

local AbhiError = {}
AbhiError.__index = AbhiError


function AbhiError.new(code, msg, ctx)
  local self = setmetatable({}, AbhiError)
  self.is_sdk_error = true
  self.sdk = "Abhi"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function AbhiError:error()
  return self.msg
end


function AbhiError:__tostring()
  return self.msg
end


return AbhiError

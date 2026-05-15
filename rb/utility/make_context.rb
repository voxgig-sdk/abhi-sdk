# Abhi SDK utility: make_context
require_relative '../core/context'
module AbhiUtilities
  MakeContext = ->(ctxmap, basectx) {
    AbhiContext.new(ctxmap, basectx)
  }
end

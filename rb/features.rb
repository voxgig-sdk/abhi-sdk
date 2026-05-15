# Abhi SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module AbhiFeatures
  def self.make_feature(name)
    case name
    when "base"
      AbhiBaseFeature.new
    when "test"
      AbhiTestFeature.new
    else
      AbhiBaseFeature.new
    end
  end
end

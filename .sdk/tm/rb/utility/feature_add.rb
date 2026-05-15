# Abhi SDK utility: feature_add
module AbhiUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end

# Abhi SDK feature factory

from feature.base_feature import AbhiBaseFeature
from feature.test_feature import AbhiTestFeature


def _make_feature(name):
    features = {
        "base": lambda: AbhiBaseFeature(),
        "test": lambda: AbhiTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()

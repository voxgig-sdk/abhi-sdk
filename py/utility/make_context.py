# Abhi SDK utility: make_context

from core.context import AbhiContext


def make_context_util(ctxmap, basectx):
    return AbhiContext(ctxmap, basectx)

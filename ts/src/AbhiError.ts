
import { Context } from './Context'


class AbhiError extends Error {

  isAbhiError = true

  sdk = 'Abhi'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  AbhiError
}


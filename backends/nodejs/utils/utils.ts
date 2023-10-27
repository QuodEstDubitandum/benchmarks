export interface Context {
  n: number
  array?: number[]
  cores?: number
  object?: {
    a: string
    b: string
    c: string
    d: string
    e: string
    f: string
    g: string
    h: string
    i: string
    j: string
    k: string
    l: string
    m: string
    n: string
    o: string
    p: string
    q: string
  }
  deserializeString?: string
  binaryImage?: Buffer
}

export async function measurePerformance(
  f: (ctx: Context) => void | Promise<void>,
  ctx: Context,
  isAsync: boolean
) {
  const start = process.hrtime()
  for (let i = 0; i < 5; i++) {
    isAsync ? await f(ctx) : f(ctx)
  }
  const diff = process.hrtime(start)
  const timeInMs = (diff[0] * 1e3 + diff[1] * 1e-6) / 5
  return timeInMs
}

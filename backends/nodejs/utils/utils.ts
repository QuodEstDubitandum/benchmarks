export interface Context {
  n: number
  quickSortArray?: number[]
}

export function measurePerformance(f: (ctx: Context) => void, ctx: Context): number {
  const start = process.hrtime()
  for (let i = 0; i < 5; i++) {
    f(ctx)
  }
  const diff = process.hrtime(start)
  const timeInMs = (diff[0] * 1e3 + diff[1] * 1e-6) / 5
  return timeInMs
}

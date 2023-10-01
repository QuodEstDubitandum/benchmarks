import express from 'express'
import { Context, measurePerformance } from '../utils/utils'

const router = express.Router()

// router
router.get('/prime', async function (req, res) {
  const time = await measurePerformance(prime, { n: 10000000 }, false)
  res.send(time.toFixed(0))
})
router.get('/fast-fibonacci', async function (req, res) {
  const time = await measurePerformance(fastFibonacci, { n: 1000000 }, false)
  res.send(time.toFixed(0))
})
router.get('/quicksort', async function (req, res) {
  const n = 1000000
  const arr = Array(n)
  for (let i = 0; i < arr.length; i++) {
    arr[i] = n - i
  }
  const time = await measurePerformance(quicksort, { n: 10000000, array: arr }, false)
  res.send(time.toFixed(0))
})
router.get('/two-sum', async function (req, res) {
  const n = 1000000
  const arr = Array(n)
  for (let i = 0; i < arr.length; i++) {
    arr[i] = i + 1
  }
  const time = await measurePerformance(twoSum, { n: n, array: arr }, false)
  res.send(time.toFixed(0))
})

// functions

// /prime
export function prime(ctx: Context) {
  const n = ctx.n

  let p = 2
  const primes = Array(n + 1).fill(false)

  let j = n - 1

  while (p * p <= n) {
    if (!primes[p - 1]) {
      for (let i = p * p; i <= n; i += p) {
        if (!primes[i - 1]) {
          j -= 1
        }
        primes[i - 1] = true
      }
    }
    p += 1
  }
  console.log('Number of prime numbers until ' + n + ': ' + j)
}

// /fast-fibonacci
function fastFibonacci(ctx: Context) {
  const n = ctx.n
  const baseMatrix = [
    [BigInt(1), BigInt(1)],
    [BigInt(1), BigInt(0)]
  ]
  powerMatrix(baseMatrix, n)
}

function powerMatrix(matrix: bigint[][], n: number): bigint[][] {
  if (n === 1) {
    return matrix
  }
  if (n % 2) {
    return matrixMult(matrix, powerMatrix(matrix, n - 1))
  }
  const root = powerMatrix(matrix, n / 2)
  return matrixMult(root, root)
}

function matrixMult(a: bigint[][], b: bigint[][]): bigint[][] {
  const res: bigint[][] = [[], []]

  res[0][0] = a[0][0] * b[0][0] + a[0][1] * b[1][0]
  res[0][1] = a[0][0] * b[0][1] + a[0][1] * b[1][1]
  res[1][0] = a[1][0] * b[0][0] + a[1][1] * b[1][0]
  res[1][1] = a[1][0] * b[0][1] + a[1][1] * b[1][1]

  return res
}

// /quicksort
function quicksort(ctx: Context) {
  const baseArray = ctx.array as number[]

  console.log(divideArray(baseArray, 0, baseArray.length - 1)[0])
}

function divideArray(array: number[], left: number, right: number) {
  if (left >= right) {
    return array
  }

  const pivot = Math.floor((right + left) / 2)

  let temp = array[pivot]
  array[pivot] = array[right]
  array[right] = temp

  let j = left
  for (let i = left; i < right + 1; i++) {
    if (array[i] < array[right]) {
      temp = array[i]
      array[i] = array[j]
      array[j] = temp
      j++
    }
  }

  temp = array[j]
  array[j] = array[right]
  array[right] = temp

  divideArray(array, left, j - 1)
  divideArray(array, j + 1, right)

  return array
}

// /two-sum
function twoSum(ctx: Context) {
  const n = ctx.n
  const array = ctx.array as number[]
  let count = 0
  let complement
  const map = new Map<number, number>()

  for (const number of array) {
    complement = n - number
    if (map.has(complement)) {
      count += map.get(complement)!
    }

    if (map.has(number)) {
      const oldNumber = map.get(number)!
      map.set(number, oldNumber + 1)
    }
    if (!map.has(number)) {
      map.set(number, 1)
    }
  }
  console.log(count)
}

export default router

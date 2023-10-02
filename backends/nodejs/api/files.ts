import express from 'express'
import os from 'os'
import { Context, measurePerformance } from '../utils/utils'
import fs from 'fs'

const cores = os.cpus().length

const router = express.Router()

// router
router.get('/write', async function (req, res) {
  const time = await measurePerformance(writeFileSync, { n: 10000, cores: cores }, false)
  res.send(time.toFixed(0))
})
router.get('/write-parallel', async function (req, res) {
  const time = await measurePerformance(writeFileParallel, { n: 10000, cores: cores }, true)
  res.send(time.toFixed(0))
})
router.get('/read', async function (req, res) {
  // creating files
  const n = 1000000
  for (let i = 0; i < cores; i++) {
    const file = fs.openSync(`test${i}.txt`, 'a')
    const string = Array(n).fill('Gotta go fast').join('\n')
    fs.writeSync(file, string)
    fs.closeSync(file)
  }
  // measuring performance of read
  const time = await measurePerformance(readFileSync, { n: n, cores: cores }, false)
  // deleting files
  for (let i = 0; i < cores; i++) {
    fs.unlinkSync(`test${i}.txt`)
  }
  res.send(time.toFixed(0))
})
router.get('/read-parallel', async function (req, res) {
  // creating files
  const n = 1000000
  for (let i = 0; i < cores; i++) {
    const file = fs.openSync(`test${i}.txt`, 'a')
    const string = Array(n).fill('Gotta go fast').join('\n')
    fs.writeSync(file, string)
    fs.closeSync(file)
  }
  // measuring performance of read
  const time = await measurePerformance(readFileParallel, { n: n, cores: cores }, true)
  // deleting files
  for (let i = 0; i < cores; i++) {
    fs.unlinkSync(`test${i}.txt`)
  }
  res.send(time.toFixed(0))
})

// /write
function writeFileSync(ctx: Context) {
  for (let i = 0; i < cores; i++) {
    const file = fs.openSync('test.txt', 'a')
    for (let j = 0; j < ctx.n; j++) {
      fs.writeSync(file, 'Gotta go fast\n')
    }
    fs.closeSync(file)
    fs.unlinkSync('test.txt')
  }
}

// /write-parallel
async function writeFileParallel(ctx: Context) {
  const tasks = []

  for (let i = 0; i < cores; i++) {
    const task = (() => {
      return new Promise<void>((resolve) => {
        fs.open(`test${i}.txt`, 'a', async (err, fd) => {
          await writeToFile(fd, ctx.n)
          fs.unlink(`test${i}.txt`, () => {
            resolve()
          })
        })
      })
    })()
    tasks.push(task)
  }

  await Promise.all(tasks)
}

async function writeToFile(fd: number, n: number) {
  for (let j = 0; j < n; j++) {
    fs.write(fd, 'Gotta go fast\n', () => {})
  }
  fs.close(fd)
}

function readFileSync(ctx: Context) {
  for (let i = 0; i < ctx.cores!; i++) {
    fs.readFileSync(`test${i}.txt`)
  }
}

async function readFileParallel(ctx: Context) {
  const tasks = []
  for (let i = 0; i < ctx.cores!; i++) {
    const task = (() => {
      return new Promise<void>((resolve) => {
        fs.readFile(`test${i}.txt`, () => {
          resolve()
        })
      })
    })()
    tasks.push(task)
  }
  await Promise.all(tasks)
}

export default router

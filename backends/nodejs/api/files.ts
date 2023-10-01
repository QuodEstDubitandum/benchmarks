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
    const task = (async () => {
      fs.open(`test${i}.txt`, 'w', async function (err, fd) {
        await writeToFile(fd, ctx.n)
        fs.unlink(`test${i}.txt`, () => {})
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

export default router

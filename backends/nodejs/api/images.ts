import express from 'express'
import { Context, measurePerformance } from '../utils/utils'
import sharp from 'sharp'

const router = express.Router()

router.get('/decode', async function (req, res) {
  const time = await measurePerformance(decodeImage, { n: 10 }, true)
  res.send(time.toFixed(0))
})

router.get('/encode', async function (req, res) {
  const imgData = await sharp('../../public/sample.jpg').raw().toBuffer()
  const time = await measurePerformance(encodeImage, { n: 1, binaryImage: imgData }, true)
  res.send(time.toFixed(0))
})

// /decode
async function decodeImage(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    await sharp('../../public/sample.jpg').raw().toBuffer()
  }
}

// /encode
async function encodeImage(ctx: Context) {
  await sharp(ctx.binaryImage, {
    raw: {
      width: 1000,
      height: 563,
      channels: 3 // assuming RGB format
    }
  }).toFile('../../public/sample.png')
}

export default router

import express from 'express'
import { Context, measurePerformance } from '../utils/utils'

const router = express.Router()

// router
router.get('/serialize', async function (req, res) {
  const sampleString =
    'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.'
  const time = await measurePerformance(
    serializeJson,
    {
      n: 10000,
      object: {
        a: sampleString,
        b: sampleString,
        c: sampleString,
        d: sampleString,
        e: sampleString,
        f: sampleString,
        g: sampleString,
        h: sampleString,
        i: sampleString,
        j: sampleString,
        k: sampleString,
        l: sampleString,
        m: sampleString,
        n: sampleString,
        o: sampleString,
        p: sampleString,
        q: sampleString
      }
    },
    false
  )
  res.send(time.toFixed(0))
})

router.post('/deserialize', async function (req, res) {
  let data = ''
  let time = 0
  req.on('data', (chunk) => {
    data += chunk
  })
  req.on('end', async () => {
    time = await measurePerformance(deserializeJson, { n: 10000, deserializeString: data }, false)
    res.send(time.toFixed(0))
  })
})

// /serialize
function serializeJson(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    JSON.stringify(ctx.object)
  }
}

// /deserialize
function deserializeJson(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    JSON.parse(ctx.deserializeString as string)
  }
}
export default router

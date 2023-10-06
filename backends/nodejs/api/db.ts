import express from 'express'
import { Context, measurePerformance } from '../utils/utils'
import Database from 'better-sqlite3'

const router = express.Router()

const db = new Database('../../sqlite.db')
db.exec('PRAGMA journal_mode=DELETE;')
// db.pragma('journal_mode = WAL')

// router
router.get('/select', async function (req, res) {
  const time = await measurePerformance(select, { n: 10 }, false)
  res.send(time.toFixed(0))
})
router.get('/insert', async function (req, res) {
  const time = await measurePerformance(insert, { n: 10 }, false)
  db.exec(
    `
        DELETE FROM electric_cars
        WHERE vin = 'test' AND county = 'test'
    `
  )

  res.send(time.toFixed(0))
})
router.get('/update', async function (req, res) {
  const time = await measurePerformance(update, { n: 10 }, false)
  res.send(time.toFixed(0))
})
router.get('/delete', async function (req, res) {
  const time = del({ n: 50 })
  res.send(time.toFixed(0))
})

// functions

// /select
function select(ctx: Context) {
  let row: any
  for (let i = 0; i < ctx.n; i++) {
    row = db
      .prepare(
        `
			SELECT COUNT(*) as count FROM electric_cars 
			WHERE vehicletype = ? AND model = ?
		`
      )
      .get('BEV', 'MODEL 3')
  }
  console.log(row.count)
}

// /insert
function insert(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    db.exec(
      `
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES ('test', 'test', 'test', 'test', 1, 1, 'test', 'test', 'test', 1, 'test')
		`
    )
  }
}

// /update
function update() {
  db.exec(
    `
    UPDATE electric_cars
		SET model = 'I4'
		WHERE model = 'I3'
  `
  )
  db.exec(
    `
    UPDATE electric_cars
		SET model = 'I3'
		WHERE model = 'I4'
  `
  )
}

// /delete
function del(ctx: Context) {
  let time = 0
  let start: [number, number]
  let diff: [number, number]
  for (let i = 0; i < ctx.n; i++) {
    db.exec(
      `
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES ('test', 'test', 'test', 'test', 1, 1, 'test', 'test', 'test', 1, 'test')
		`
    )
    start = process.hrtime()
    db.exec(
      `
			DELETE FROM electric_cars
			WHERE vin = 'test' AND county = 'test'
		`
    )
    diff = process.hrtime(start)
    time += (diff[0] * 1e3 + diff[1] * 1e-6) / 5
  }
  return time
}
export default router

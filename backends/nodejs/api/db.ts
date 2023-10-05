import express from 'express'
import { Context, measurePerformance } from '../utils/utils'
import Database from 'better-sqlite3'

const router = express.Router()

const db = new Database('../../sqlite.db')
// db.exec('PRAGMA journal_mode=DELETE;')
db.pragma('journal_mode = WAL')

// router
router.get('/select', async function (req, res) {
  const time = await measurePerformance(select, { n: 10 }, false)
  res.send(time.toFixed(0))
})
router.get('/insert', async function (req, res) {
  const time = await measurePerformance(insert, { n: 1000 }, false)
  db.prepare(
    `
        DELETE FROM electric_cars
        WHERE vin = 'test' AND county = 'test'
    `
  ).run()

  res.send(time.toFixed(0))
})

// functions

// /select
function select(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    db.prepare(
      `
			SELECT COUNT(*) as count FROM electric_cars 
			WHERE vehicletype = ? AND model = ?
		`
    ).get('BEV', 'MODEL 3')
  }
}

// /insert
function insert(ctx: Context) {
  for (let i = 0; i < ctx.n; i++) {
    db.prepare(
      `
			INSERT INTO electric_cars (vin, county, city, state, postalCode, modelYear, make, model, vehicleType, electricRange, vehicleLocation)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
    ).run('test', 'test', 'test', 'test', 1, 1, 'test', 'test', 'test', 1, 'test')
  }
}
export default router

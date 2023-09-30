import express from 'express'
import algorithmRoutes from './api/algorithms'

const app = express()

app.use('/algorithms', algorithmRoutes)

app.listen(3002, () => {
  console.log('Server is running on port 3002')
})

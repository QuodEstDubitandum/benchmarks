import express from 'express'
import algorithmRoutes from './api/algorithms'
import filesRoutes from './api/files'

const app = express()

app.use('/algorithms', algorithmRoutes)
app.use('/files', filesRoutes)

app.listen(3002, () => {
  console.log('Server is running on port 3002')
})

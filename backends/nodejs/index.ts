import express from 'express'
import algorithmRoutes from './api/algorithms'
import filesRoutes from './api/files'
import dbRoutes from './api/db'
import jsonRoutes from './api/json'

const app = express()

app.use('/algorithms', algorithmRoutes)
app.use('/files', filesRoutes)
app.use('/db', dbRoutes)
app.use('/json', jsonRoutes)

app.listen(3002, () => {
  console.log('Server is running on port 3002')
})

import { Router } from 'express'

import axios from '../../plugins/axios'

const router = Router()

router.get('/messages', async (req, res) => {
  const params = {limit: req.query.limit}
  if (req.query.start_id) {
    params.start_id = req.query.start_id
  } else {
    params.end_id = req.query.end_id
  }
  let { data } = await axios.get(`http://localhost:1323/messages`, {params})
  res.json(data)
})

router.post('/messages', async (req, res) => {
  let { data } = await axios.post(`http://localhost:1323/messages`, {
    text: req.body.text
  }, {
    headers: { Authorization: `Bearer ${req.session.auth.jwt}` }
  })
  res.json(data)
})

export default router

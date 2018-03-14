import { Router } from 'express'

import axios from '../../plugins/axios'

const router = Router()

router.post('/login', async (req, res) => {
  try {
    let { data } = await axios.post(`http://localhost:1323/login`, {
      name: req.body.name,
      password: req.body.password
    })
    req.session.auth = {
      user: data.user,
      jwt: data.jwt
    }
    res.json(req.session.auth.user)
  } catch (e) {
    res.status(401).json({ error: 'Bad credentials' })
  }
})

router.post('/signup', async (req, res) => {
  try {
    let { data } = await axios.post(`http://localhost:1323/signup`, {
      name: req.body.name,
      password: req.body.password
    })
    console.log(data)
    req.session.auth = {
      user: data.user,
      jwt: data.jwt
    }
    res.json(req.session.auth.user)
  } catch (e) {
    res.status(400).json({ error: 'Bad request' })
  }
})

router.post('/logout', (req, res) => {
  delete req.session.auth
  res.json({ ok: true })
})

export default router

import { Router } from 'express'

import messages from './messages'
import user from './user'

const router = Router()

router.use(messages)
router.use(user)

export default router

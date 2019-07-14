import {
  QUEUE_TIMED_OUT_REQUEST,
  RETRY_TIMED_OUT_REQUEST,
  CLEAR_TIMED_OUT_REQUEST
} from 'actions/types'

export const queueTimedOutRequest = request => (
  {
      type: QUEUE_TIMED_OUT_REQUEST,
      request
  }
)

export const retryTimedOutRequest = requestId => (
  {
    type: RETRY_TIMED_OUT_REQUEST,
    id: requestId
  }
)

export const clearTimedOutRequest = () => (
  {
    type: CLEAR_TIMED_OUT_REQUEST
  }
)

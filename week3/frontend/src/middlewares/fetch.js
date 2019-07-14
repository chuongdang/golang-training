import { showError } from 'helpers/notification'
import { logOut } from 'actions/user_action'

import NavigationService from 'helpers/navigation'
import { timeoutDuration, REQUEST_TIMEOUT_CODE } from 'constants/timeout_constant'
import uuidv4 from 'uuid/v4'
import { queueTimedOutRequest } from 'actions/fetch_action'
import { QUEUE_TIMED_OUT_REQUEST } from 'actions/types'

const timeOutRejectPromise = timeout => (new Promise((resolve, reject) => {
  setTimeout(reject, timeout, REQUEST_TIMEOUT_CODE)
}))


const fetchMiddleware = store => next => (action) => {
  if (!action || !action.request || action.type === QUEUE_TIMED_OUT_REQUEST) {
    return next(action)
  }

  const fetchAction = { ...action }

  if (!fetchAction.id) {
    fetchAction.id = uuidv4()
  }

  const timeout = fetchAction.timeout || timeoutDuration.default
  const fetchRequest = createFetchRequest(fetchAction, store.getState().user.token)

  Promise.race([fetchRequest, timeOutRejectPromise(timeout)])
    .then((response) => {
      console.log(response)
      if (!response.ok || response.status >= 400) {
        if (response.status === 403) {
          store.dispatch(logOut())
          NavigationService.resetNavigation()
        }
        if (response.statusText) {
          throw Error(response.statusText)
        }
      }
      return response.json()
    })
    .then((data) => {
      if (!data) {
        //throw Error(message)
      }
      store.dispatch({
        type: `${action.type}_OK`,
        params: action.params,
        data
      })
    })
    .catch((error) => {
      console.log(error)
      if (error === REQUEST_TIMEOUT_CODE) {
        store.dispatch(queueTimedOutRequest(fetchAction))
      } else {
        showError(error.Error)
      }
      store.dispatch({
        type: `${action.type}_ERR`,
        params: action.params,
        data: error
      })
    })

  return next(action)
}

/* Helpers */

const createFetchRequest = (action, token) => {
  const {
    url = '/',
    method = 'GET',
    headers = {},
  } = action.request


  const body = ['POST'].indexOf(method) >= 0 && action.params
    ? getEncodedUrlParams(action.params)
    : undefined

  const queryParams = ['GET'].indexOf(method) >= 0 && action.params
    ? `?${getEncodedUrlParams(action.params)}`
    : ''
  const TOKEN_KEY = 'X-ACCESS-TOKEN'

  const fetchRequest = fetch(`${getApiUrl(url)}${queryParams}`, {
    method,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
      [TOKEN_KEY]: token,
      ...headers
    },
    body
  })
  return fetchRequest
}

const getApiUrl = (url) => {
  if (/^http[s]*:\/\/.+/.test(url)) {
    return url
  }
  return `http://127.0.0.1:8989/${url}`
}

const getEncodedUrlParams = params => Object.keys(params)
  .map(key => (params[key] ? `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}` : ''))
  .join('&')


export default fetchMiddleware

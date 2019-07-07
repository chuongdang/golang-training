import { combineReducers } from 'redux'
import {
  LOAD_PRODUCT,
  LOAD_PRODUCT_OK
} from 'actions/types'

const DEFAULT_STATE = {
  loading: false,
  products: []
}

const product = (state = DEFAULT_STATE, { type, ...payload }) => {
  switch (type) {
    case LOAD_PRODUCT:
      return {
        ...state,
        loading: true
      }
    case LOAD_PRODUCT_OK:
        return {
          ...state,
          loading: false,
          products: payload.data
        }
    default:
      return state
  }
}

const DEFAULT_USER_STATE = {
  token: 'token'
}

const user = (state = DEFAULT_USER_STATE, action) => {
  switch (action.type) {
    default:
      return state
  }
}

const nordicApp = combineReducers({
  product,
  user
})

export default nordicApp

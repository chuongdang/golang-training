import { combineReducers } from 'redux'
import {
  INIT
} from '../actions'

function todos(state = [], action) {
  switch (action.type) {
    case INIT:
      return [
        ...state,
        {
          products: action.,
          completed: false
        }
      ]
    case TOGGLE_TODO:
      return state.map((todo, index) => {
        if (index === action.index) {
          return Object.assign({}, todo, {
            completed: !todo.completed
          })
        }
        return todo
      })
    default:
      return state
  }
}

const todoApp = combineReducers({
  visibilityFilter,
  todos
})

export default todoApp

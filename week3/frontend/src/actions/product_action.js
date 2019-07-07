import {
  LOAD_PRODUCT
} from 'actions/types'

export const loadProducts = ({ page, limit }) => ({
  type: LOAD_PRODUCT,
  params: { page, limit },
  request: {
    url: 'products',
    method: 'GET'
  }
})

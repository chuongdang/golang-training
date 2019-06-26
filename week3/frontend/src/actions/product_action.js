/*
 * action types
 */

export const INIT = 'INIT'

/*
 * action creators
 */

export function init(limit) {
  return { type: INIT, limit }
}

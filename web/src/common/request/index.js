import axios from "axios"

export const apiGet = (url = '', param = {}) => axios.get(getFinalUrl(url), param)

export const apiPost = (url = '', param = {}) => axios.post(getFinalUrl(url), param)

function getFinalUrl(url) {
  const retUrl = url.indexOf('/') === 0 ? url : `/${url}`
  return ENV_IS_PRO ? retUrl : 'test' + retUrl
}
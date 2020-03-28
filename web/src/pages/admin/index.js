import './index.less';
import { apiPost } from 'src/common/request/index';
import $ from 'jquery'

const btn = $('.btn')
let isSubmiting = false

btn.click(() => {
  if (isSubmiting) return

  const title = document.querySelector('.title-input').value;
  const content = document.querySelector( '.content-input').innerText;

  btn.toggleClass()
  isSubmiting = true
  
  apiPost('/api/article', {title, content })
    .then(() => {
      isSubmiting = false
    })
    .catch((err) => {
      isSubmiting = false
      console.error(err)
      alert('上传出错')
    })
})


import './index.less'

import axios from 'axios'

document.querySelector('.btn').addEventListener('click', () => {
  const title = document.querySelector('.title-input').value
  const content = document.querySelector('.content-input').innerText

  axios.post('/api/article', {title, content })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  })
})


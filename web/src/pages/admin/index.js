import './index.less';
import { apiGet } from 'src/common/request/index';

document.querySelector('.btn').addEventListener('click', async () => {
  try {
    const title = document.querySelector('.title-input').value;
    const content = document.querySelector( '.content-input').innerText;

    await apiGet('/api/article', {title, content })

    alert('上传成功～')
  } catch (error) {
    console.error(error)
    alert('上传失败')
  }
});

